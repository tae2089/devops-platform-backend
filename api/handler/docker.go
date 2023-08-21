package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/usecase"
)

type DockerHandler struct {
	DockerUsecase usecase.DockerUsecase
}

func (d *DockerHandler) GetDockerFile(c *gin.Context) {
	err := d.DockerUsecase.GetDockerFile(c.Request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
