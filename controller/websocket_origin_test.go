package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebsocketOriginAllowed(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		origin    string
		allowlist string
		want      bool
	}{
		{"non browser client", "gateway.example.com", "", "", true},
		{"same origin", "gateway.example.com", "https://gateway.example.com", "", true},
		{"trusted console", "gateway.example.com", "https://console.example.com", "https://console.example.com", true},
		{"untrusted origin", "gateway.example.com", "https://evil.example.com", "https://console.example.com", false},
		{"wildcard does not weaken websocket checks", "gateway.example.com", "https://evil.example.com", "*", false},
		{"invalid origin", "gateway.example.com", "javascript:alert(1)", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("CORS_ALLOWED_ORIGINS", tt.allowlist)
			request := httptest.NewRequest("GET", "https://"+tt.host+"/v1/realtime", nil)
			request.Host = tt.host
			if tt.origin != "" {
				request.Header.Set("Origin", tt.origin)
			}
			assert.Equal(t, tt.want, websocketOriginAllowed(request))
		})
	}
}
