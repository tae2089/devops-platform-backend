package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/domain"
	"net/http"
)

type ServiceHandler struct{}

func (s *ServiceHandler) RegisterService(c *gin.Context) {
	requestService := &domain.RequestService{}
	if err := c.ShouldBindJSON(requestService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
