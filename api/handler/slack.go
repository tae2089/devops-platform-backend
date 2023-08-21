package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

type SlackHandler struct {
	SlackUtil slack.Util
}

func (s *SlackHandler) CallBack(c *gin.Context) {
	payload := c.PostForm("payload")
	i, err := s.SlackUtil.GetCallbackPayload(&payload)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	switch i.View.PrivateMetadata {
	case "/slash":
		log.Println("check slash")
		log.Println(i.User.ID)
	case "/front-deploy":
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
