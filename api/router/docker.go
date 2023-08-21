package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/api/handler"
	"github.com/tae2089/devops-platform-backend/usecase"
	"github.com/tae2089/devops-platform-backend/util/docker"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

func newDockerRouter(timeout time.Duration, dockerRouter *gin.RouterGroup, slackUtil slack.Util, dockerUtil docker.Util) {
	usecase := usecase.NewDockerUsecase(slackUtil, dockerUtil)
	dockerHandler := handler.DockerHandler{
		DockerUsecase: usecase,
	}
	dockerRouter.POST("/file", dockerHandler.GetDockerFile)
}
