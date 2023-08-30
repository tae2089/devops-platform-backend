package usecase

import (
	"github.com/slack-go/slack"
	"github.com/tae2089/devops-platform-backend/repository"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
	slackUtil "github.com/tae2089/devops-platform-backend/util/slack"
)

type SlackUsecase interface {
	GetCallbackPayload(payload *string) (*slack.InteractionCallback, error)
	RegistJenkinsFrontJob(callback *slack.InteractionCallback) error
	RegistJenkinsProject(callback *slack.InteractionCallback) error
}

func NewSlackUsecase(slackUtil slackUtil.Util, jenkinsUtil jenkins.Util, githubUtil github.Util, userRepository repository.UserRepository, jenkinsRepository repository.JenkinsRepository) SlackUsecase {
	return &slackUsecase{
		slackUtil:         slackUtil,
		jenkinsUtil:       jenkinsUtil,
		githubUtil:        githubUtil,
		userRepository:    userRepository,
		jenkinsRepository: jenkinsRepository,
	}
}
