package usecase

import (
	"github.com/tae2089/devops-platform-backend/util/docker"
	"github.com/tae2089/devops-platform-backend/util/slack"
	"strings"
)

var _ (DockerUsecase) = (*dockerUsecase)(nil)

type dockerUsecase struct {
	slackUtil  slack.Util
	dockerUtil docker.Util
}

func (d dockerUsecase) GetDockerFile(lang string) string {
	lang = strings.ToLower(lang)
	file := d.dockerUtil.GetFileByLang(lang)
	return file
}
