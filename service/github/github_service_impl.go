package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"github.com/tae2089/devops-platform-backend/domain"
	"github.com/tae2089/devops-platform-backend/exception"
)

type gitServiceImpl struct {
	client *github.Client
}

var _ GitService = (*gitServiceImpl)(nil)

// GetHooksForRepo functions to output hooks stored in the repo.
func (g gitServiceImpl) GetHooksForRepo(ctx context.Context, hookDto domain.RequestGithubWebhookDto) ([]*github.Hook, error) {
	//TODO implement me
	hooks, _, err := g.client.Repositories.ListHooks(ctx, hookDto.Owner, hookDto.Repo, &github.ListOptions{})
	if err != nil {
		return nil, err
	}
	return hooks, nil
}

// RegisterWebhookForJenkins is a function that registers a webhook in a GitHub repository for  Jenkins. Since we're using HTTPS, insecure_ssl is set to false.
// Also, if at least one webhook is already registered, registration will fail.
func (g gitServiceImpl) RegisterWebhookForJenkins(ctx context.Context, hookDto domain.RequestGithubWebhookDto) error {
	_, err := g.isExistWebhook(ctx, hookDto.Owner, hookDto.Repo)
	if err != nil {
		return err
	}
	hookURL := fmt.Sprintf("%s/generic-webhook-trigger/invoke?token=%s", hookDto.TargetUrl, hookDto.Token)
	_, _, err = g.client.Repositories.CreateHook(ctx, hookDto.Owner, hookDto.Repo, &github.Hook{
		Name:   github.String(hookDto.Repo),
		Active: github.Bool(true),
		Events: []string{
			"push",
		},
		Config: map[string]interface{}{
			"content_type": "json",
			"url":          github.String(hookURL),
			"insecure_ssl": github.Bool(false),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// ModifyWebhookForJenkins is used to change the URL of a webhook. The data required in the request is the same as RegisterWebhookForJenkins.
func (g gitServiceImpl) ModifyWebhookForJenkins(ctx context.Context, hookDto domain.RequestGithubWebhookDto) error {
	id, err := g.isExistWebhook(ctx, hookDto.Owner, hookDto.Repo)
	if id == nil {
		if err != nil {
			return err
		}
		return exception.NotFoundHooks
	}
	hook, _, err := g.client.Repositories.GetHook(ctx, hookDto.Owner, hookDto.Repo, *id)
	if err != nil {
		return err
	}
	hook.URL = github.String(hookDto.TargetUrl)
	_, _, err = g.client.Repositories.EditHook(ctx, hookDto.Owner, hookDto.Repo, *id, hook)
	if err != nil {
		return err
	}
	return nil
}

// isExistWebhook checks to see if the hook exists in the repository, and if so, returns an error and the hook's id.
// If not, it returns nil. If you're updating a hook, you'll need the id value and use it when returning.
// If you're registering a hook, you just need to check for an error.
func (g gitServiceImpl) isExistWebhook(ctx context.Context, owner, repoName string) (*int64, error) {
	hooks, _, err := g.client.Repositories.ListHooks(ctx, owner, repoName, nil)
	if err != nil {
		return nil, err
	}
	if len(hooks) != 0 {
		return hooks[0].ID, exception.AlreadyHookExistsError
	}
	// webhook not exist
	return nil, nil
}
