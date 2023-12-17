package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/api/handler"
	"github.com/tae2089/devops-platform-backend/usecase"
	"github.com/tae2089/devops-platform-backend/util/docker"
)

func newDockerRouter(timeout time.Duration, dockerRouter *gin.RouterGroup, dockerUtil docker.Util) {
	usecase := usecase.NewDockerUsecase(dockerUtil)
	dockerHandler := handler.DockerHandler{
		DockerUsecase: usecase,
	}
	dockerRouter.GET("/file", dockerHandler.GetDockerFile)
}
