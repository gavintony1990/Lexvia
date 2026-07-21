package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performCORSPreflight(t *testing.T, origin string) *httptest.ResponseRecorder {
	t.Helper()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CORS())
	router.OPTIONS("/v1/chat/completions", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	request := httptest.NewRequest(http.MethodOptions, "/v1/chat/completions", nil)
	request.Header.Set("Origin", origin)
	request.Header.Set("Access-Control-Request-Method", http.MethodPost)
	request.Header.Set("Access-Control-Request-Headers", "authorization,content-type")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

func TestCORSDefaultsToCredentiallessPublicRelay(t *testing.T) {
	t.Setenv("CORS_ALLOWED_ORIGINS", "")
	t.Setenv("CORS_ALLOW_CREDENTIALS", "false")

	response := performCORSPreflight(t, "https://client.example")

	assert.Equal(t, "*", response.Header().Get("Access-Control-Allow-Origin"))
	assert.Empty(t, response.Header().Get("Access-Control-Allow-Credentials"))
}

func TestCORSCredentialsRequireExplicitAllowlist(t *testing.T) {
	t.Setenv("CORS_ALLOWED_ORIGINS", "https://console.example")
	t.Setenv("CORS_ALLOW_CREDENTIALS", "true")

	response := performCORSPreflight(t, "https://console.example")

	assert.Equal(t, "https://console.example", response.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", response.Header().Get("Access-Control-Allow-Credentials"))
}

func TestCORSCredentialsRejectWildcardConfiguration(t *testing.T) {
	t.Setenv("CORS_ALLOWED_ORIGINS", "")
	t.Setenv("CORS_ALLOW_CREDENTIALS", "true")

	assert.Panics(t, func() {
		_ = CORS()
	})
}
