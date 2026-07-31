package circuitbreaker

import (
	"net/http"
	"os"
	"strconv"
)

// RoundTripper wraps an http.RoundTripper with circuit breaker per base URL.
type RoundTripper struct {
	next http.RoundTripper
}

// NewRoundTripper creates a new circuit breaker RoundTripper.
// It wraps the given next transport with per-host circuit breaker logic.
func NewRoundTripper(next http.RoundTripper) *RoundTripper {
	if next == nil {
		next = http.DefaultTransport
	}
	return &RoundTripper{next: next}
}

// RoundTrip implements http.RoundTripper.
// It uses the request's host (or a configured channel base URL) as the circuit breaker key.
func (t *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Use the base URL from the request context if set, otherwise use Host header.
	key := req.URL.Host
	if key == "" {
		key = req.URL.String()
	}

	var resp *http.Response
	err := ExecuteChannelRequest(key, func() error {
		var err error
		resp, err = t.next.RoundTrip(req)
		return err
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Enabled returns whether the circuit breaker is enabled via env var.
// Default: enabled (true). Set CIRCUIT_BREAKER_ENABLED=false to disable.
func Enabled() bool {
	v := os.Getenv("CIRCUIT_BREAKER_ENABLED")
	if v == "" {
		return true
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return true
	}
	return b
}

// WrapTransport wraps the given transport with a circuit breaker RoundTripper.
func WrapTransport(transport http.RoundTripper) http.RoundTripper {
	if !Enabled() {
		return transport
	}
	return NewRoundTripper(transport)
}

// WrapClient wraps the given HTTP client's transport with circuit breaker.
func WrapClient(client *http.Client) *http.Client {
	if !Enabled() || client == nil {
		return client
	}

	transport := client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	// Clone the client to avoid mutating the original
	clone := &http.Client{}
	*clone = *client
	clone.Transport = WrapTransport(transport)
	return clone
}