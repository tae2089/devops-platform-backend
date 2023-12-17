package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/ent"
	"github.com/tae2089/devops-platform-backend/repository"
	"github.com/tae2089/devops-platform-backend/usecase"
	"github.com/tae2089/devops-platform-backend/util/aws"
	"github.com/tae2089/devops-platform-backend/util/docker"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
)

func SetUp(client *ent.Client, timeout time.Duration, g *gin.Engine) {

	awsUtils := make(map[string]aws.Util)

	for _, profile := range config.GetAwsConfig().Profiles {
		awsUtils[profile] = aws.NewAwsUtil(profile)
	}
	// utils setup
	jenkinsUtil := jenkins.NewJenkinsUtil(config.GetJenkinsConfig())
	githubUtil := github.NewGithubUtil(config.GetGithubConfig())
	dockerUtil := docker.NewDockerUtil()

	// repository setup

	userRepository := repository.NewUserRepository(client)
	jenkinsRepository := repository.NewJenkinsRepository(client)

	// jenkins router setup
	jenkinsRouter := g.Group("/jenkins")
	//jenkinsRouter.Use(middleware.VerifySlack())
	jenkinsUsecase := usecase.NewJenkinsUsecase(jenkinsUtil, githubUtil, userRepository, jenkinsRepository)
	newJenkinsRouter(timeout, jenkinsRouter, jenkinsUsecase)

	// health router setup
	healthRouter := g.Group("/")
	newHealthRouter(timeout, healthRouter)

	//docker router setup
	dockerRouter := g.Group("/docker")
	//dockerRouter.Use(middleware.VerifySlack())
	newDockerRouter(timeout, dockerRouter, dockerUtil)
}
