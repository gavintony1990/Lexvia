package service

import (
	"net/http"

	"github.com/gavintony1990/Lexvia/service/httputil"
)

// GetHttpClient returns the singleton HTTP client with circuit breaker.
// Deprecated: use httputil.GetHttpClient() instead.
func GetHttpClient() *http.Client {
	return httputil.GetHttpClient()
}

// GetHttpClientWithProxy returns an HTTP client configured with the given proxy URL.
// Deprecated: use httputil.GetHttpClientWithProxy() instead.
func GetHttpClientWithProxy(proxyURL string) (*http.Client, error) {
	return httputil.GetHttpClientWithProxy(proxyURL)
}

// NewProxyHttpClient creates a new HTTP client with proxy support.
// Deprecated: use httputil.NewProxyHttpClient() instead.
func NewProxyHttpClient(proxyURL string) (*http.Client, error) {
	return httputil.NewProxyHttpClient(proxyURL)
}

// ResetProxyClientCache clears all cached proxy HTTP clients.
// Deprecated: use httputil.ResetProxyClientCache() instead.
func ResetProxyClientCache() {
	httputil.ResetProxyClientCache()
}

// InitHttpClient initializes the default HTTP client with circuit breaker.
// Deprecated: use httputil.InitHttpClient() instead.
func InitHttpClient() {
	httputil.InitHttpClient()
}