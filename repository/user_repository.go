package repository

import (
	"context"

	"github.com/tae2089/devops-platform-backend/ent"
	"github.com/tae2089/devops-platform-backend/ent/user"
)

type UserRepository interface {
	FindBySlackID(ctx context.Context, slackID string) (*ent.User, error)
	NewTransaction(ctx context.Context) (*ent.Tx, error)
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepositoryImpl{
		client,
	}
}

type userRepositoryImpl struct {
	client *ent.Client
}

// FindBySlackID implements UserRepository.
func (u *userRepositoryImpl) FindBySlackID(ctx context.Context, slackID string) (*ent.User, error) {
	return u.client.User.Query().Where(user.SlackIDEQ(slackID)).First(ctx)
}

func (u *userRepositoryImpl) NewTransaction(ctx context.Context) (*ent.Tx, error) {
	return u.client.Tx(ctx)
}
