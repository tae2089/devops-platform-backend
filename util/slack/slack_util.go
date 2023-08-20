package slack

import (
	"net/http"
	"sync"

	"github.com/slack-go/slack"
	"github.com/tae2089/devops-platform-backend/config"
)

type SlackUtil interface {
	OpenView(triggerId string, modalRequest slack.ModalViewRequest) error
	SlashCommandParse(request *http.Request) (string, error)
	GenerateModalRequest() slack.ModalViewRequest
	GetUserProfile(userId string) (string, error)
	GetUsersRealName(userId ...string) ([]string, error)
}

var (
	slackOnce sync.Once
	slackUtil SlackUtil
)

func NewSlackUtil(slackConfig *config.SlackBot) SlackUtil {
	client := slack.New(slackConfig.AccessToken)
	if slackUtil == nil {
		slackOnce.Do(func() {
			slackUtil = &slackUtilImpl{
				client,
			}
		})
	}
	return slackUtil
}
