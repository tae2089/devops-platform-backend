package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/api/handler"
	"github.com/tae2089/devops-platform-backend/usecase"
)

func newJenkinsRouter(timeout time.Duration, group *gin.RouterGroup, jenkinsUsecase usecase.JenkinsUsecase) {
	jenkinsHandler := &handler.JenkinsHandler{
		JenkinsUsecase: jenkinsUsecase,
	}
	group.POST("/create/job", jenkinsHandler.CheckHealth)
	// group.GET("/healthz", healthRouter.CheckHealth)
}
