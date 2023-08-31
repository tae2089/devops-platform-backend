package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	goSlack "github.com/slack-go/slack"
	"github.com/tae2089/devops-platform-backend/domain"
	"github.com/tae2089/devops-platform-backend/ent"
	"github.com/tae2089/devops-platform-backend/exception"
	"github.com/tae2089/devops-platform-backend/repository"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

var _ (JenkinsUsecase) = (*jenkinsUsecaseImpl)(nil)

type jenkinsUsecaseImpl struct {
	slackUtil         slack.Util
	jenkinsUtil       jenkins.Util
	githubUtil        github.Util
	userRepository    repository.UserRepository
	jenkinsRepository repository.JenkinsRepository
}

// RegisteProject implements JenkinsUsecase.
func (j *jenkinsUsecaseImpl) RegistProject(request *http.Request) error {
	slashCommand, err := j.slackUtil.GetSlashCommandParse(request)
	if err != nil {
		return err
	}
	user, err := j.userRepository.FindBySlackID(context.Background(), slashCommand.UserID)

	if err != nil {
		return exception.IsEntityNotFound(err)
	}

	if err = exception.IsAdminRole(user.Roles); err != nil {
		return err
	}

	var modalRequest goSlack.ModalViewRequest = goSlack.ModalViewRequest{}
	err = j.slackUtil.OpenView(slashCommand.TriggerID, modalRequest)
	if err != nil {
		fmt.Printf("Error opening view: %s", err)
		return err
	}
	return nil
}

// RegistFrontJob implements JenkinsUsecase.
func (j *jenkinsUsecaseImpl) RegistJob(request *http.Request) error {
	ctx := context.Background()
	slashCommand, err := j.slackUtil.GetSlashCommandParse(request)
	if err != nil {
		return err
	}
	var modalRequest goSlack.ModalViewRequest = goSlack.ModalViewRequest{}
	switch slashCommand.Text {
	case "front":
		//TODO: necessary using db connection
		projects, err := j.jenkinsRepository.FindAll(ctx)
		if err != nil {
			return err
		}
		selectOptions := getSelectOptions(projects)
		modalRequest = j.slackUtil.GenerateFrontDeployModal(selectOptions...)
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

func getSelectOptions(projects []*ent.JenkinsProject) []domain.SelectOption {
	selectOptions := make([]domain.SelectOption, len(projects))
	for _, project := range projects {
		selectOptions = append(selectOptions, domain.SelectOption{
			Text:  project.ProjectName,
			Value: project.ProjectValue,
		})
	}
	return selectOptions
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
