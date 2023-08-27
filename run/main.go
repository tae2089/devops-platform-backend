package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/bob-logging/logger"
	router "github.com/tae2089/devops-platform-backend/api/router"
	config "github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/ent"
	"go.uber.org/zap"
)

func main() {
	client, err := ent.Open("postgres", config.GetDsn())
	if err != nil {
		logger.Error(err, zap.String("message", "failed opening connection to mysql"))
	}
	defer client.Close()
	timeout := time.Duration(3) * time.Second
	gin := gin.Default()
	router.SetUp(client, timeout, gin)
	gin.Run(":8080")
}
