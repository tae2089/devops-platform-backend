package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/api/middleware"
	"github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/ent"
	"github.com/tae2089/devops-platform-backend/usecase"
	"github.com/tae2089/devops-platform-backend/util/aws"
	"github.com/tae2089/devops-platform-backend/util/docker"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

func SetUp(client *ent.Client, timeout time.Duration, g *gin.Engine) {

	awsUtils := make(map[string]aws.Util)

	for _, profile := range config.GetAwsConfig().Profiles {
		awsUtils[profile] = aws.NewAwsUtil(profile)
	}

	jenkinsUtil := jenkins.NewJenkinsUtil(config.GetJenkinsConfig())
	githubUtil := github.NewGithubUtil(config.GetGithubConfig())
	slackUtil := slack.NewSlackUtil(config.GetSlackBotConfig())
	dockerUtil := docker.NewDockerUtil()

	// jenkins router setup
	jenkinsRouter := g.Group("/jenkins")
	jenkinsRouter.Use(middleware.VerifySlack())
	jenkinsUsecase := usecase.NewJenkinsUsecase(slackUtil, jenkinsUtil, githubUtil)
	newJenkinsRouter(timeout, jenkinsRouter, jenkinsUsecase)

	// slack router setup

	slackRouter := g.Group("/slack")
	slackRouter.Use(middleware.VerifySlack())
	slackUsecase := usecase.NewSlackUsecase(slackUtil, jenkinsUtil, githubUtil)
	newSlackRouter(timeout, slackRouter, slackUsecase)

	// health router setup
	healthRouter := g.Group("/")
	newHealthRouter(timeout, healthRouter)

	//docker router setup
	dockerRouter := g.Group("/docker")
	dockerRouter.Use(middleware.VerifySlack())
	newDockerRouter(timeout, dockerRouter, slackUtil, dockerUtil)
}
