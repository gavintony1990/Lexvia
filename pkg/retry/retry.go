package retry

import (
	"context"
	"math"
	"math/rand"
	"time"
)

// Config configures the exponential backoff retry mechanism.
type Config struct {
	// MaxRetries is the maximum number of retry attempts (excluding the initial attempt).
	MaxRetries int
	// InitialBackoff is the wait time before the first retry.
	InitialBackoff time.Duration
	// MaxBackoff caps the total wait time between retries.
	MaxBackoff time.Duration
	// Multiplier is the exponential factor applied to backoff after each retry.
	Multiplier float64
	// JitterFactor adds randomness to backoff: actual delay = backoff * (1 ± jitterFactor).
	JitterFactor float64
}

// DefaultConfig returns a sensible default retry configuration.
func DefaultConfig() Config {
	return Config{
		MaxRetries:     3,
		InitialBackoff: 500 * time.Millisecond,
		MaxBackoff:     10 * time.Second,
		Multiplier:     2.0,
		JitterFactor:   0.2,
	}
}

// IsRetryable returns true if the error should be retried.
// Override this function to customize retryable error detection.
type IsRetryable func(err error) bool

// AlwaysRetryable returns true for all errors (default).
func AlwaysRetryable(err error) bool {
	return err != nil
}

// DoWithRetry executes fn with exponential backoff retry.
// fn should return (shouldRetry bool, err error).
// If fn returns shouldRetry=false, the error is considered non-retryable and is returned immediately.
// If fn returns shouldRetry=true, the operation is retried up to MaxRetries times.
func DoWithRetry(ctx context.Context, cfg Config, isRetryable IsRetryable, fn func() error) error {
	if isRetryable == nil {
		isRetryable = AlwaysRetryable
	}

	var err error
	backoff := cfg.InitialBackoff

	for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
		// Execute the operation
		err = fn()

		// Success — return nil
		if err == nil {
			return nil
		}

		// Non-retryable error — return immediately
		if !isRetryable(err) {
			return err
		}

		// Last attempt — return the error
		if attempt >= cfg.MaxRetries {
			return err
		}

		// Wait with exponential backoff + jitter
		delay := calculateBackoff(backoff, cfg)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}

		// Exponential increase for next iteration
		backoff = time.Duration(float64(backoff) * cfg.Multiplier)
		if backoff > cfg.MaxBackoff {
			backoff = cfg.MaxBackoff
		}
	}

	return err
}

// calculateBackoff applies jitter to the base backoff duration.
func calculateBackoff(base time.Duration, cfg Config) time.Duration {
	if cfg.JitterFactor <= 0 {
		return base
	}

	jitter := time.Duration(float64(base) * cfg.JitterFactor * (2*rand.Float64() - 1))
	delay := base + jitter
	if delay < 0 {
		delay = 0
	}
	return delay
}

// IntPow computes base^exp for use in backoff calculation.
func IntPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

// FullJitterBackoff calculates backoff with full jitter: [0, base*multiplier^attempt)
// This is the "full jitter" strategy from AWS's exponential backoff article.
func FullJitterBackoff(attempt int, base time.Duration, maxBackoff time.Duration, multiplier float64) time.Duration {
	expBackoff := float64(base) * math.Pow(multiplier, float64(attempt))
	if expBackoff > float64(maxBackoff) {
		expBackoff = float64(maxBackoff)
	}
	return time.Duration(rand.Float64() * expBackoff)
}