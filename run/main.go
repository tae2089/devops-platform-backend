package main

import (
	"time"

	"github.com/gin-gonic/gin"
	router "github.com/tae2089/devops-platform-backend/api/router"
	_ "github.com/tae2089/devops-platform-backend/config"
)

func main() {
	timeout := time.Duration(3) * time.Second
	gin := gin.Default()
	router.SetUp(timeout, gin)
	gin.Run(":8080")
}
