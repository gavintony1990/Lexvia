package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	allowedOrigins := splitAndTrim(os.Getenv("CORS_ALLOWED_ORIGINS"))
	allowCredentials := common.GetEnvOrDefaultBool("CORS_ALLOW_CREDENTIALS", false)
	if len(allowedOrigins) == 0 {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = allowedOrigins
	}
	if allowCredentials && config.AllowAllOrigins {
		panic("CORS_ALLOW_CREDENTIALS requires an explicit CORS_ALLOWED_ORIGINS allowlist")
	}
	config.AllowCredentials = allowCredentials
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{
		"Authorization",
		"Content-Type",
		"New-Api-User",
		"X-Api-Key",
		"Anthropic-Version",
		"Anthropic-Beta",
		"X-Goog-Api-Key",
		"OpenAI-Beta",
		"X-Request-Id",
	}
	config.AllowHeaders = append(config.AllowHeaders, splitAndTrim(os.Getenv("CORS_ADDITIONAL_ALLOWED_HEADERS"))...)
	return cors.New(config)
}

func splitAndTrim(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	values := make([]string, 0)
	for _, item := range strings.Split(value, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		if strings.ContainsAny(item, "\r\n") {
			panic(fmt.Sprintf("invalid CORS configuration value %q", item))
		}
		values = append(values, item)
	}
	return values
}

func Version() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-New-Api-Version", common.Version)
		c.Next()
	}
}
