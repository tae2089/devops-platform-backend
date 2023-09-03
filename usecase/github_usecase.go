package usecase

import (
	"net/http"

	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

type GithubUsecase interface {
	RegisterWebhook(request *http.Request) error
}

func NewGithubUsecase(slackUtil slack.Util, githubUtil github.Util) GithubUsecase {
	return &githubUsecaseImpl{
		slackUtil: slackUtil,
		// githubUtil: githubUtil,
	}
}
