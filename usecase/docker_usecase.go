package usecase

import (
	"github.com/tae2089/devops-platform-backend/util/docker"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

type DockerUsecase interface {
	GetDockerFile(imageName string) string
}

func NewDockerUsecase(slackUtil slack.Util, dockerUtil docker.Util) DockerUsecase {
	return &dockerUsecase{
		slackUtil,
		dockerUtil,
	}
}
