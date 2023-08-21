package usecase

import (
	"log"
	"net/http"

	"github.com/tae2089/devops-platform-backend/util/docker"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

type dockerUsecase struct {
	slackUtil  slack.Util
	dockerUtil docker.Util
}

// GetDockerFile implements DockerUsecase.
func (d *dockerUsecase) GetDockerFile(request *http.Request) error {
	slackCommand, err := d.slackUtil.GetSlashCommandParse(request)
	if err != nil {
		return err
	}
	templateData := ""
	switch slackCommand.Text {
	case "java":
		templateData = d.dockerUtil.GetJava()
		break
	case "go":
		templateData = d.dockerUtil.GetGolang()
		break
	case "js":
		templateData = d.dockerUtil.GetJavaScript()
		break
	default:
		templateData = "NONE"
		break
	}
	log.Println(templateData)
	blocks := d.slackUtil.GetDockerCodeBlocks(templateData)
	err = d.slackUtil.PostMessageWithBlocks(slackCommand.ChannelID, blocks)
	if err != nil {
		return err
	}
	return nil
}

var _ (DockerUsecase) = (*dockerUsecase)(nil)
