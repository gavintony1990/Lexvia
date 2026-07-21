package router

import (
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAPIRouterDoesNotExposeSubscriptionManagement(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	SetApiRouter(engine)

	for _, route := range engine.Routes() {
		path := strings.ToLower(route.Path)
		if strings.Contains(path, "/subscription") {
			t.Fatalf("subscription management route is still exposed: %s %s", route.Method, route.Path)
		}
	}
}
