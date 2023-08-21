package usecase

import (
	"net/http"

	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

type JenkinsUsecase interface {
	RegisterLunchPayment(request *http.Request) error
}

func NewJenkinsUsecase(slackUtil slack.Util, jenkinsUtil jenkins.JenkinsUtil, githubUtil github.GithubUtil) JenkinsUsecase {
	return &jenkinsUsecaseImpl{
		slackUtil,
		jenkinsUtil,
		githubUtil,
	}
}
