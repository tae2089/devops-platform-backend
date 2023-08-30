package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/devops-platform-backend/usecase"
	"go.uber.org/zap"
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
		logger.Info("check slash", zap.String("userID", i.User.ID))
	case "/front-deploy":
		err = s.SlackUsecase.RegistJenkinsFrontJob(i)
	case "/jenkins-project":
		err = s.SlackUsecase.RegistJenkinsProject(i)
	default:
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err != nil {
		logger.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusAccepted, "")
}
