package domain

import (
	"github.com/tae2089/devops-platform-backend/domain/deploy"
	"github.com/tae2089/devops-platform-backend/domain/language"
	"github.com/tae2089/devops-platform-backend/domain/lifecycle"
	"github.com/tae2089/devops-platform-backend/domain/tier"
)

type RequestService struct {
	Name                string         `json:"name" binding:"required"`
	UserID              string         `json:"user_id" binding:"required"`
	GitRepoURL          string         `json:"git_repo_url" binding:"required"`
	Language            language.Type  `json:"language,omitempty"`
	Lifecycle           lifecycle.Type `json:"lifecycle,omitempty"`
	Tier                tier.Type      `json:"tier,omitempty"`
	JenkinsURL          string         `json:"jenkins_url,omitempty"`
	IsUsingGithubAction bool           `json:"is_using_github_action,omitempty"`
	IsUsingGitlabRunner bool           `json:"is_using_gitlab_runner,omitempty"`
	MonitoringURL       string         `json:"monitoring_url,omitempty"`
	DeployRepoURL       string         `json:"deploy_repo_url,omitempty"`
	VaultURL            string         `json:"vault_url,omitempty"`
	DeployKind          deploy.Kind    `json:"deploy_kind,omitempty"`
	SlackChannelURL     string         `json:"slack_channel_url,omitempty"`
}
