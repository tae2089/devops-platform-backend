package route

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetUp(timeout time.Duration, g *gin.Engine) {

	healthRouter := g.Group("/")
	newHealthRouter(timeout, healthRouter)
}
