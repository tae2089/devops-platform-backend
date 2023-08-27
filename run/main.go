package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	router "github.com/tae2089/devops-platform-backend/api/router"
	config "github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/ent"
)

func main() {
	client, err := ent.Open("postgres", config.GetDsn())
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	timeout := time.Duration(3) * time.Second
	gin := gin.Default()
	router.SetUp(timeout, gin)
	gin.Run(":8080")
}
