package service

import (
	"net/http"

	"github.com/gavintony1990/Lexvia/service/httputil"
	"github.com/gin-gonic/gin"
)

// CloseResponseBodyGracefully closes the response body and logs any error.
// Deprecated: use httputil.CloseResponseBodyGracefully() instead.
func CloseResponseBodyGracefully(httpResponse *http.Response) {
	httputil.CloseResponseBodyGracefully(httpResponse)
}

// ShouldCopyUpstreamHeader checks whether a given upstream response header
// should be copied to the client response.
// Deprecated: use httputil.ShouldCopyUpstreamHeader() instead.
func ShouldCopyUpstreamHeader(c *gin.Context, k string, v []string) bool {
	return httputil.ShouldCopyUpstreamHeader(c, k, v)
}

// IOCopyBytesGracefully copies data from the upstream response to the client.
// Deprecated: use httputil.IOCopyBytesGracefully() instead.
func IOCopyBytesGracefully(c *gin.Context, src *http.Response, data []byte) {
	httputil.IOCopyBytesGracefully(c, src, data)
}