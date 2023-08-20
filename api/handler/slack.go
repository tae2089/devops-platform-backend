package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	util "github.com/tae2089/devops-platform-backend/util/slack"
)

type SlackHandler struct {
	SlackUtil util.SlackUtil
}

func (s *SlackHandler) CallBack(c *gin.Context) {
	var i slack.InteractionCallback

	err := json.Unmarshal([]byte(c.PostForm("payload")), &i)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	switch i.View.PrivateMetadata {
	case "/slash":
		log.Println("check slash")
		log.Println(i.User.ID)
	default:
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusAccepted, "")
}
