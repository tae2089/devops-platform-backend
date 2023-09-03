package slack

import (
	"net/http"
	"sync"

	"github.com/slack-go/slack"
	"github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/domain"
)

type Util interface {
	OpenView(triggerId string, modalRequest slack.ModalViewRequest) error
	SlashCommandParse(request *http.Request) (string, error)
	GenerateModalRequest() slack.ModalViewRequest
	GetUserProfile(userId string) (string, error)
	GetUsersRealName(userId ...string) ([]string, error)
	GetSlashCommandParse(request *http.Request) (slack.SlashCommand, error)
	PostMessageWithBlocks(channelId string, blocks []slack.Block) error
	GetDockerCodeBlocks(content string) []slack.Block
	GetCallbackPayload(payload *string) (*slack.InteractionCallback, error)
	GenerateFrontDeployModal(options ...domain.SelectOption) slack.ModalViewRequest
	GetJenkinsJobResultBlocks(content string) []slack.Block
	GenerateProjectRegisterModal() slack.ModalViewRequest
	GenerateGithubWebhookModal() slack.ModalViewRequest
}

var (
	slackOnce sync.Once
	util      Util
)

func NewSlackUtil(slackConfig *config.SlackBot) Util {
	client := slack.New(slackConfig.AccessToken)
	if util == nil {
		slackOnce.Do(func() {
			util = &slackUtil{
				client,
			}
		})
	}
	return util
}
