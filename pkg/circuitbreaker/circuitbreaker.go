package circuitbreaker

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

// State represents the circuit breaker state.
type State int32

const (
	StateClosed   State = iota // Normal operation
	StateOpen                  // Failures exceeded threshold, requests are rejected immediately
	StateHalfOpen              // Probing: allow one request to test if upstream recovered
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateOpen:
		return "open"
	case StateHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

// Config configures a circuit breaker instance.
type Config struct {
	// FailureThreshold is the number of consecutive failures required to open the circuit.
	FailureThreshold int
	// SuccessThreshold is the number of consecutive successes in half-open state to close the circuit.
	SuccessThreshold int
	// HalfOpenTimeout is how long to wait before transitioning from open to half-open.
	HalfOpenTimeout time.Duration
	// MaxRequests is the maximum number of requests allowed through while half-open.
	MaxRequests int
}

// DefaultConfig returns a sensible default configuration.
func DefaultConfig() Config {
	return Config{
		FailureThreshold: 5,
		SuccessThreshold: 2,
		HalfOpenTimeout:  30 * time.Second,
		MaxRequests:      1,
	}
}

// Breaker is a circuit breaker keyed by a string identifier (e.g., channel base URL).
type Breaker struct {
	mu     sync.Mutex
	cfg    Config
	state  State

	failures         int
	successes        int
	halfOpenRequests int
	lastStateChange  time.Time
}

// NewBreaker creates a new circuit breaker with the given config.
func NewBreaker(cfg Config) *Breaker {
	return &Breaker{
		cfg:   cfg,
		state: StateClosed,
	}
}

// ErrOpen is returned when the circuit is open and the request is rejected.
var ErrOpen = errors.New("circuit breaker is open")

// Execute runs fn if the circuit is closed or half-open.
// If the circuit is open, it returns ErrOpen immediately without calling fn.
// fn should return an error indicating whether the operation succeeded.
// If fn returns nil, the operation is considered a success.
func (b *Breaker) Execute(fn func() error) error {
	if !b.ready() {
		return ErrOpen
	}

	err := fn()
	b.record(err)
	return err
}

// ready returns true if the circuit allows a request through.
func (b *Breaker) ready() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	switch b.state {
	case StateClosed:
		return true
	case StateOpen:
		if time.Since(b.lastStateChange) >= b.cfg.HalfOpenTimeout {
			b.setState(StateHalfOpen)
			b.halfOpenRequests = 0
			return true
		}
		return false
	case StateHalfOpen:
		if b.halfOpenRequests < b.cfg.MaxRequests {
			b.halfOpenRequests++
			return true
		}
		return false
	default:
		return false
	}
}

// record updates the circuit breaker state based on the result of an operation.
func (b *Breaker) record(err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if err == nil {
		b.onSuccess()
	} else {
		b.onFailure()
	}
}

func (b *Breaker) onSuccess() {
	switch b.state {
	case StateClosed:
		b.failures = 0
	case StateHalfOpen:
		b.successes++
		if b.successes >= b.cfg.SuccessThreshold {
			b.setState(StateClosed)
			b.failures = 0
			b.successes = 0
		}
	}
}

func (b *Breaker) onFailure() {
	switch b.state {
	case StateClosed:
		b.failures++
		if b.failures >= b.cfg.FailureThreshold {
			b.setState(StateOpen)
		}
	case StateHalfOpen:
		b.setState(StateOpen)
		b.successes = 0
	}
}

func (b *Breaker) setState(s State) {
	b.state = s
	b.lastStateChange = time.Now()
}

// State returns the current state of the breaker.
func (b *Breaker) State() State {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.state
}

// Reset resets the breaker to the closed state.
func (b *Breaker) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.state = StateClosed
	b.failures = 0
	b.successes = 0
	b.halfOpenRequests = 0
	b.lastStateChange = time.Now()
}

// ---------------------------------------------------------------------------
// BreakerGroup — a collection of breakers keyed by a string (e.g., channel URL)
// ---------------------------------------------------------------------------

// BreakerGroup manages circuit breakers keyed by an identifier.
type BreakerGroup struct {
	mu       sync.RWMutex
	breakers map[string]*Breaker
	cfg      Config
}

// NewBreakerGroup creates a new breaker group with the given config.
func NewBreakerGroup(cfg Config) *BreakerGroup {
	return &BreakerGroup{
		breakers: make(map[string]*Breaker),
		cfg:      cfg,
	}
}

// Get returns the breaker for the given key, creating it if necessary.
func (g *BreakerGroup) Get(key string) *Breaker {
	g.mu.RLock()
	b, ok := g.breakers[key]
	g.mu.RUnlock()
	if ok {
		return b
	}

	g.mu.Lock()
	defer g.mu.Unlock()
	// Double-check after acquiring write lock
	if b, ok := g.breakers[key]; ok {
		return b
	}
	b = NewBreaker(g.cfg)
	g.breakers[key] = b
	return b
}

// Delete removes the breaker for the given key (e.g., when a channel is deleted).
func (g *BreakerGroup) Delete(key string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.breakers, key)
}

// ResetAll resets all breakers to the closed state.
func (g *BreakerGroup) ResetAll() {
	g.mu.RLock()
	defer g.mu.RUnlock()
	for _, b := range g.breakers {
		b.Reset()
	}
}

// Count returns the number of active breakers.
func (g *BreakerGroup) Count() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.breakers)
}

// ---------------------------------------------------------------------------
// Global breaker group for upstream channels
// ---------------------------------------------------------------------------

var (
	// ChannelBreakers is the global breaker group keyed by channel base URL.
	ChannelBreakers = NewBreakerGroup(DefaultConfig())

	// ErrBreakerOpen is the error returned when a channel's circuit breaker is open.
	ErrBreakerOpen = errors.New("channel circuit breaker is open, rejecting request")
)

// ExecuteChannelRequest runs fn for the given channel base URL.
// If the circuit breaker is open, it returns ErrBreakerOpen immediately.
func ExecuteChannelRequest(channelBaseURL string, fn func() error) error {
	breaker := ChannelBreakers.Get(channelBaseURL)
	return breaker.Execute(fn)
}

// ResetChannelBreaker resets the breaker for a specific channel.
func ResetChannelBreaker(channelBaseURL string) {
	ChannelBreakers.Delete(channelBaseURL)
}

// ResetAllChannelBreakers resets all channel breakers.
func ResetAllChannelBreakers() {
	ChannelBreakers.ResetAll()
}

// ChannelBreakerStats holds the current state of a channel's circuit breaker.
type ChannelBreakerStats struct {
	ChannelURL string `json:"channel_url"`
	State      string `json:"state"`
}

// AllChannelBreakerStats returns the state of all active channel breakers.
func AllChannelBreakerStats() []ChannelBreakerStats {
	ChannelBreakers.mu.RLock()
	defer ChannelBreakers.mu.RUnlock()

	stats := make([]ChannelBreakerStats, 0, len(ChannelBreakers.breakers))
	for url, breaker := range ChannelBreakers.breakers {
		stats = append(stats, ChannelBreakerStats{
			ChannelURL: url,
			State:      breaker.State().String(),
		})
	}
	return stats
}

// Ensure BreakerGroup exports are accessible for the stats function.
// We use the internal fields directly in AllChannelBreakerStats, which is fine
// since it's in the same package.

// Ensure atomic operations are used (for future use).
var _ atomic.Int32