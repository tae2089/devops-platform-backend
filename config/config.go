package config

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/tae2089/bob-logging/logger"
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

type Aws struct {
	Profiles []string `env:"PROFILES" envSeparator:";"`
}

type DB struct {
	Port     string `env:"DB_PORT,required"`
	Host     string `env:"DB_HOST,required"`
	DBname   string `env:"DB_NAME,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
}

var (
	githubConfig   = &Github{}
	jenkinsConfig  = &Jenkins{}
	slackBotConfig = &SlackBot{}
	awsConfig      = &Aws{}
	dbConfig       = &DB{}
)

func init() {

	var err error

	if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "" {
		projectDir := getProjectDir() + "/.env"
		err = godotenv.Load(projectDir)
	}

	if err != nil {
		logger.Error(err)
		return
	}

	if err = env.Parse(githubConfig); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	if err = env.Parse(jenkinsConfig); err != nil {
		panic(err)
	}

	if err = env.Parse(slackBotConfig); err != nil {
		panic(err)
	}

	if err = env.Parse(awsConfig); err != nil {
		panic(err)
	}
	if err = env.Parse(dbConfig); err != nil {
		panic(err)
	}
	// logger.Info("check data", zap.Any("awsconfig", awsConfig))
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

func GetAwsConfig() *Aws {
	return awsConfig
}

func GetDBConfig() *DB {
	return dbConfig
}

func GetDsn() string {
	// return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBname) // MYSQL
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.DBname, dbConfig.Password)
}
