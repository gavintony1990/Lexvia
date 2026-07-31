package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealthz K8s liveness probe — 只检查进程是否存活
func GetHealthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// GetReadyz K8s readiness probe — 检查服务是否可处理请求
func GetReadyz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}
