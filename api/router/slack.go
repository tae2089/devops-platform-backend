package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/api/handler"
	"github.com/tae2089/devops-platform-backend/usecase"
)

func newSlackRouter(timeout time.Duration, slackRouter *gin.RouterGroup, slackUsecase usecase.SlackUsecase) {
	slackHandler := &handler.SlackHandler{
		SlackUsecase: slackUsecase,
	}
	slackRouter.POST("/callback", slackHandler.CallBack)
}
