package usecase

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"
	"github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/domain"
	"github.com/tae2089/devops-platform-backend/repository"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
	slackUtil "github.com/tae2089/devops-platform-backend/util/slack"
)

var _ (SlackUsecase) = (*slackUsecase)(nil)

type slackUsecase struct {
	slackUtil         slackUtil.Util
	githubUtil        github.Util
	jenkinsUtil       jenkins.Util
	userRepository    repository.UserRepository
	jenkinsRepository repository.JenkinsRepository
}

// RegistGithubWebhook implements SlackUsecase.
func (s *slackUsecase) RegistGithubWebhook(callback *slack.InteractionCallback) error {
	ctx := context.Background()
	owner := callback.View.State.Values["owner"]["owner"]
	repositoryName := callback.View.State.Values["repositoryName"]["repositoryName"]
	jenkinsToken := callback.View.State.Values["jenkinsToken"]["jenkinsToken"]
	jenkinsConfig := config.GetJenkinsConfig()

	err := s.githubUtil.RegisterWebhookForJenkins(ctx, &domain.RequestGithubWebhookDto{
		Owner:     owner.Value,
		Repo:      repositoryName.Value,
		Token:     jenkinsToken.Value,
		TargetUrl: jenkinsConfig.URL,
	})
	if err != nil {
		return err
	}

	return nil
}

// RegistJenkinsProject implements SlackUsecase.
func (s *slackUsecase) RegistJenkinsProject(callback *slack.InteractionCallback) error {
	ctx := context.Background()
	projectName := callback.View.State.Values["projectName"]["projectName"]
	projectValue := callback.View.State.Values["projectValue"]["projectValue"]
	err := s.jenkinsRepository.SaveJenkinsProject(ctx, projectName.Value, projectValue.Value)
	if err != nil {
		return err
	}
	channelID := callback.Channel.ID
	if callback.Channel.Name == "" {
		channelID = callback.User.ID
	}
	resultBlocks := s.slackUtil.GetJenkinsJobResultBlocks(fmt.Sprintf("success project registerd - %s", projectName.Value))
	err = s.slackUtil.PostMessageWithBlocks(channelID, resultBlocks)
	if err != nil {
		return err
	}
	return nil
}

// RegistJenkinsFrontJob implements SlackUsecase.
func (s *slackUsecase) RegistJenkinsFrontJob(callback *slack.InteractionCallback) error {
	projectSelect := callback.View.State.Values["project selector"]["project_select"]
	repositoryField := callback.View.State.Values["repository block"]["repository"]
	branchField := callback.View.State.Values["branch block"]["branch"]
	webhookField := callback.View.State.Values["webhook block"]["webhook"]
	jenkinsFileField := callback.View.State.Values["jenkinsfile block"]["jenkinsfile"]
	jenkinsJobField := callback.View.State.Values["jenkinsjob block"]["jenkinsjob"]
	content := s.jenkinsUtil.GetJenkinsJobContent(repositoryField.Value, branchField.Value, projectSelect.SelectedOption.Value, webhookField.Value, jenkinsFileField.Value)
	_, err := s.jenkinsUtil.CreateJob(&jenkinsJobField.Value, &projectSelect.SelectedOption.Text.Text, &content)
	if err != nil {
		return err
	}
	responseJob := &domain.ResultMessageJenkinsJob{
		Project:      projectSelect.SelectedOption.Value,
		WebhookToken: webhookField.Value,
		GitURL:       repositoryField.Value,
		Branch:       branchField.Value,
		FileName:     jenkinsFileField.Value,
	}
	channelID := callback.Channel.ID
	if callback.Channel.Name == "" {
		channelID = callback.User.ID
	}

	resultBlocks := s.slackUtil.GetJenkinsJobResultBlocks(responseJob.Write())
	err = s.slackUtil.PostMessageWithBlocks(channelID, resultBlocks)
	if err != nil {
		return err
	}
	return nil
}

// GetCallbackPayload implements SlackUsecase.
func (s *slackUsecase) GetCallbackPayload(payload *string) (*slack.InteractionCallback, error) {
	i, err := s.slackUtil.GetCallbackPayload(payload)
	if err != nil {
		return nil, err
	}
	return i, nil
}
