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
	lang := c.Query("lang")
	file := d.DockerUsecase.GetDockerFile(lang)
	if file == "" {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//c.JSON(http.StatusOK, gin.H{})
	c.Data(http.StatusOK, "text/plain", []byte(file))
}
