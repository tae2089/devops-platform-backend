package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Github struct {
	Token string `env:"GITHUB_TOKEN,required"`
}

type Jenkins struct {
	URL      string `env:"JENKINS_URL,required"`
	User     string `env:"JENKINS_USER,required"`
	PassWord string `env:"JENKINS_PASSWORD,required"`
}

type SlackBot struct {
	AccessToken string `env:"SLACK_BOT_ACCESS_TOKEN,required"`
	SecretToken string `env:"SLACK_BOT_SECRET_TOKEN,required"`
}

var (
	githubConfig   = &Github{}
	jenkinsConfig  = &Jenkins{}
	slackBotConfig = &SlackBot{}
)

func init() {

	var err error

	if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "" {
		projectDir := getProjectDir() + "/.env"
		err = godotenv.Load(projectDir)
	}

	if err != nil {
		log.Println("Error loading.env file")
		return
	}

	if err := env.Parse(githubConfig); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	if err := env.Parse(jenkinsConfig); err != nil {
		panic(err)
	}

	if err := env.Parse(slackBotConfig); err != nil {
		panic(err)
	}

}

func getProjectDir() string {
	projectDir := ""
	_, filename, _, _ := runtime.Caller(0)
	projectDir = path.Join(path.Dir(filename), "..")
	return projectDir
}

func GetGithubConfig() *Github {
	return githubConfig
}

func GetJenkinsConfig() *Jenkins {
	return jenkinsConfig
}

func GetSlackBotConfig() *SlackBot {
	return slackBotConfig
}
