package usecase

import (
	"net/http"

	"github.com/tae2089/devops-platform-backend/repository"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
)

type JenkinsUsecase interface {
	RegistJob(request *http.Request) error
	RegistProject(request *http.Request) error
}

func NewJenkinsUsecase(jenkinsUtil jenkins.Util, githubUtil github.Util, userRepository repository.UserRepository, jeninsRepository repository.JenkinsRepository) JenkinsUsecase {
	return &jenkinsUsecaseImpl{
		jenkinsUtil,
		githubUtil,
		userRepository,
		jeninsRepository,
	}
}
