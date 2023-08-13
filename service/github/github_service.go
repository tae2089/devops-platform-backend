package github

import (
	"context"
	"github.com/google/go-github/v53/github"
	"github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/domain"
	"golang.org/x/oauth2"
	"sync"
)

type GitService interface {
	RegisterWebhookforJenkins(ctx context.Context, hookDto domain.RegisterGithubWebhookDto) error
}

var (
	gitOnce    sync.Once
	gitService GitService
)

func NewGitService(gitConfig config.Github) GitService {
	if gitService == nil {
		gitOnce.Do(func() {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: gitConfig.Token},
			)
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)
			gitService = &gitServiceImpl{
				client,
			}
		})
	}
	return gitService
}
