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

var githubConfig = &Github{}

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
