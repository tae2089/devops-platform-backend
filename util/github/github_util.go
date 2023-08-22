package github

import (
	"context"
	"sync"

	"github.com/google/go-github/v53/github"
	"github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/domain"
	"golang.org/x/oauth2"
)

type Util interface {
	RegisterWebhookForJenkins(ctx context.Context, hookDto *domain.RequestGithubWebhookDto) error
	ModifyWebhookForJenkins(ctx context.Context, hookDto *domain.RequestGithubWebhookDto) error
	GetHooksForRepo(ctx context.Context, hookDto *domain.RequestGithubWebhookDto) ([]*github.Hook, error)
	UploadFile(ctx context.Context, hookDto *domain.RequestUploadFileDto) error
	DeleteWebhook(ctx context.Context, hookDto *domain.RequestGithubWebhookDto) error
}

var (
	gitOnce    sync.Once
	githubUtil Util
)

func NewGithubUtil(gitConfig *config.Github) Util {
	if githubUtil == nil {
		gitOnce.Do(func() {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: gitConfig.Token},
			)
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)
			githubUtil = &gitServiceImpl{
				client,
			}
		})
	}
	return githubUtil
}
