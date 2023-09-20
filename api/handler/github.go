package handler

import "github.com/gin-gonic/gin"

type GithubHandler struct{}

func (g *GithubHandler) RegisterWebhook(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "webhook registered",
	})
}
