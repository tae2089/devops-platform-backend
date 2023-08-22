package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/usecase"
)

type SlackHandler struct {
	SlackUsecase usecase.SlackUsecase
}

func (s *SlackHandler) CallBack(c *gin.Context) {
	payload := c.PostForm("payload")
	i, err := s.SlackUsecase.GetCallbackPayload(&payload)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	switch i.View.PrivateMetadata {
	case "/back-deploy":
		log.Println("check slash")
		log.Println(i.User.ID)
	case "/front-deploy":
		err = s.SlackUsecase.RegistJenkinsFrontJob(i)
	default:
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusAccepted, "")
}
