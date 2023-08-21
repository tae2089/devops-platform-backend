package usecase

import (
	"errors"
	"fmt"
	"net/http"

	goSlack "github.com/slack-go/slack"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

var _ (JenkinsUsecase) = (*jenkinsUsecaseImpl)(nil)

type jenkinsUsecaseImpl struct {
	slackUtil   slack.Util
	jenkinsUtil jenkins.Util
	githubUtil  github.GithubUtil
}

// RegistFrontJob implements JenkinsUsecase.
func (j *jenkinsUsecaseImpl) RegistJob(request *http.Request) error {
	slashCommand, err := j.slackUtil.GetSlashCommandParse(request)
	if err != nil {
		return err
	}
	var modalRequest goSlack.ModalViewRequest = goSlack.ModalViewRequest{}
	switch slashCommand.Text {
	case "front":
		modalRequest = j.slackUtil.GenerateFrontDeployModal("bc-labs", "pet-i")
	case "back":
		break
	default:
		return errors.New("invalid command")
	}
	err = j.slackUtil.OpenView(slashCommand.TriggerID, modalRequest)
	if err != nil {
		fmt.Printf("Error opening view: %s", err)
		return err
	}
	return nil
}

// RegisterLunchPayment implements JenkinsUsecase.
func (j *jenkinsUsecaseImpl) RegisterLunchPayment(request *http.Request) error {
	triggerId, err := j.slackUtil.SlashCommandParse(request)
	if err != nil {
		return err
	}
	modalRequest := j.slackUtil.GenerateModalRequest()
	err = j.slackUtil.OpenView(triggerId, modalRequest)
	if err != nil {
		fmt.Printf("Error opening view: %s", err)
		return err
	}
	return nil
}
