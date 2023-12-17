package usecase

import (
	"github.com/tae2089/devops-platform-backend/util/docker"
)

type DockerUsecase interface {
	GetDockerFile(imageName string) string
}

func NewDockerUsecase(dockerUtil docker.Util) DockerUsecase {
	return &dockerUsecase{
		dockerUtil,
	}
}
