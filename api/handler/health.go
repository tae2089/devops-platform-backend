package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func (h *HealthHandler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"isLive": true})
}
