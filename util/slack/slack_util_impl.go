package slack

import (
	"encoding/json"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/tae2089/bob-logging/logger"
)

var _ Util = (*slackUtil)(nil)

type slackUtil struct {
	client *slack.Client
}

// OpenProjectRegisterModal implements Util.
func (s *slackUtil) GenerateProjectRegisterModal() slack.ModalViewRequest {
	// Create a ModalViewRequest with a header and two inputs
	titleText := slack.NewTextBlockObject(slack.PlainTextType, "Front 배포", false, false)
	closeText := slack.NewTextBlockObject(slack.PlainTextType, "취소", false, false)
	submitText := slack.NewTextBlockObject(slack.PlainTextType, "제출", false, false)
	headerText := slack.NewTextBlockObject(slack.MarkdownType, "Front Web 배포하기", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	// job name 입력하기
	projectNameBlock := s.getPlainTextBlock("생성할 project name을 입력해주세요.", "project name을 입력해주세요.ex) project-name", "projectName", "projectName")
	projectValueBlock := s.getPlainTextBlock("생성할 project value를 입력해주세요.", "project value를 입력해주세요.ex) project-value", "projectValue", "projectValue")
	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			projectNameBlock,
			projectValueBlock,
		},
	}
	metadata := "/jenkins-project"
	modalRequest := s.createModalRequest(titleText, closeText, submitText, blocks, metadata)
	return modalRequest
}

func (s *slackUtil) OpenView(triggerId string, modalRequest slack.ModalViewRequest) error {
	_, err := s.client.OpenView(triggerId, modalRequest)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (s *slackUtil) postMessage(channelId string, options ...slack.MsgOption) error {
	_, _, err := s.client.PostMessage(channelId, options...)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (s *slackUtil) GetUserProfile(userId string) (string, error) {
	userProfile, err := s.client.GetUserProfile(
		&slack.GetUserProfileParameters{
			UserID: userId,
		},
	)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return userProfile.RealName, nil
}

func (s *slackUtil) GetUsersRealName(userId ...string) ([]string, error) {
	users, err := s.client.GetUsersInfo(userId...)
	userRealNameList := []string{}
	if err != nil {
		return userRealNameList, err
	}
	for _, user := range *users {
		userRealNameList = append(userRealNameList, user.RealName)
	}
	return userRealNameList, nil
}

func (s *slackUtil) SlashCommandParse(request *http.Request) (string, error) {

	slackCommand, err := slack.SlashCommandParse(request)
	if err != nil {
		return "", err
	}
	logger.Info(slackCommand.TriggerID)
	return slackCommand.TriggerID, nil
}

// GetSlashCommandParse implements SlackUtil.
func (s *slackUtil) GetSlashCommandParse(request *http.Request) (slack.SlashCommand, error) {
	slackCommand, err := slack.SlashCommandParse(request)
	if err != nil {
		return slack.SlashCommand{}, err
	}
	return slackCommand, nil
}

func (s *slackUtil) GetDockerCodeBlocks(content string) []slack.Block {
	headerText := slack.NewTextBlockObject("mrkdwn", "아래 코드를 Dockerfile에 입력해주세요.", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	codeText := slack.NewTextBlockObject("mrkdwn", "```\n"+content+"\n```", false, false)
	codeSection := slack.NewSectionBlock(codeText, nil, nil)
	return []slack.Block{
		headerSection,
		codeSection,
	}
}

func (s *slackUtil) GetJenkinsJobResultBlocks(content string) []slack.Block {
	headerText := slack.NewTextBlockObject("mrkdwn", "아래는 입력해주신 내용입니다.", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	codeText := slack.NewTextBlockObject("mrkdwn", "```\n"+content+"\n```", false, false)
	codeSection := slack.NewSectionBlock(codeText, nil, nil)
	return []slack.Block{
		headerSection,
		codeSection,
	}
}

// GetCallbackPayload implements Util.
func (*slackUtil) GetCallbackPayload(payload *string) (*slack.InteractionCallback, error) {
	var InteractionCallback slack.InteractionCallback
	err := json.Unmarshal([]byte(*payload), &InteractionCallback)
	if err != nil {
		return nil, err
	}
	return &InteractionCallback, nil
}

// PostMessageWithBlocks implements SlackUtil.
func (s *slackUtil) PostMessageWithBlocks(channelId string, blocks []slack.Block) error {
	err := s.postMessage(channelId, slack.MsgOptionBlocks(blocks...))
	if err != nil {
		return err
	}
	return nil
}
