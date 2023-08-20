package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/api/handler"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

func newSlackRouter(timeout time.Duration, slackRouter *gin.RouterGroup, slackUtil slack.SlackUtil) {
	slackHandler := &handler.SlackHandler{
		SlackUtil: slackUtil,
	}
	slackRouter.POST("/callback", slackHandler.CallBack)
}
