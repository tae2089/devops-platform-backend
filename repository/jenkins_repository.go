package repository

import (
	"context"

	"github.com/tae2089/devops-platform-backend/ent"
)

type JenkinsRepository interface {
	SaveJenkinsProjectTX(ctx context.Context, tx *ent.Tx, projectName string, projectValue string) error
	SaveJenkinsProject(ctx context.Context, projectName string, projectValue string) error
	NewTransaction(ctx context.Context) (*ent.Tx, error)
}

func NewJenkinsRepository(client *ent.Client) JenkinsRepository {
	return &jenkinsRepositoryImpl{
		client,
	}
}

type jenkinsRepositoryImpl struct {
	client *ent.Client
}

// SaveJenkinsProject implements JenkinsRepository.
func (j *jenkinsRepositoryImpl) SaveJenkinsProject(ctx context.Context, projectName string, projectValue string) error {
	_, err := j.client.JenkinsProject.Create().
		SetProjectName(projectName).
		SetProjectValue(projectValue).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// SaveJenkinsProject implements JenkinsRepository.
func (j *jenkinsRepositoryImpl) SaveJenkinsProjectTX(ctx context.Context, tx *ent.Tx, projectName string, projectValue string) error {
	_, err := tx.JenkinsProject.Create().
		SetProjectName(projectName).
		SetProjectValue(projectValue).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// NewTransaction implements JenkinsRepository.
func (j *jenkinsRepositoryImpl) NewTransaction(ctx context.Context) (*ent.Tx, error) {
	return j.client.Tx(ctx)
}