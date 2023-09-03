package usecase

import (
	"net/http"

	slackUtil "github.com/tae2089/devops-platform-backend/util/slack"
)

var _ (GithubUsecase) = (*githubUsecaseImpl)(nil)

type githubUsecaseImpl struct {
	slackUtil slackUtil.Util
}

// RegisterWebhook implements GithubUsecase.
func (g *githubUsecaseImpl) RegisterWebhook(request *http.Request) error {
	slashCommand, err := g.slackUtil.GetSlashCommandParse(request)
	if err != nil {
		return err
	}
	modalRequest := g.slackUtil.GenerateGithubWebhookModal()
	err = g.slackUtil.OpenView(slashCommand.TriggerID, modalRequest)
	if err != nil {
		return err
	}
	return nil
}
