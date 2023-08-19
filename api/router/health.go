package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/api/handler"
)

func newHealthRouter(timeout time.Duration, group *gin.RouterGroup) {
	healthRouter := handler.HealthHandler{}
	group.GET("/healthz", healthRouter.CheckHealth)
}
