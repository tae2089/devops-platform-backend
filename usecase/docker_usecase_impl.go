package usecase

import (
	"strings"

	"github.com/tae2089/devops-platform-backend/util/docker"
)

var _ (DockerUsecase) = (*dockerUsecase)(nil)

type dockerUsecase struct {
	dockerUtil docker.Util
}

func (d dockerUsecase) GetDockerFile(lang string) string {
	lang = strings.ToLower(lang)
	file := d.dockerUtil.GetFileByLang(lang)
	return file
}
