package usecase

import (
	"net/http"

	"github.com/tae2089/devops-platform-backend/repository"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

type JenkinsUsecase interface {
	RegisterLunchPayment(request *http.Request) error
	RegistJob(request *http.Request) error
	RegistProject(request *http.Request) error
}

func NewJenkinsUsecase(slackUtil slack.Util, jenkinsUtil jenkins.Util, githubUtil github.Util, userRepository repository.UserRepository, jeninsRepository repository.JenkinsRepository) JenkinsUsecase {
	return &jenkinsUsecaseImpl{
		slackUtil,
		jenkinsUtil,
		githubUtil,
		userRepository,
		jeninsRepository,
	}
}
