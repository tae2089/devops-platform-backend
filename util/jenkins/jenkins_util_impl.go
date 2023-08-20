package jenkins

import (
	"context"

	"github.com/bndr/gojenkins"
)

type jenkinsUtilImpl struct {
	jenkins *gojenkins.Jenkins
}

var _ JenkinsUtil = (*jenkinsUtilImpl)(nil)

func (j *jenkinsUtilImpl) CreateJob(jobName, folderName, content *string) (*gojenkins.Job, error) {
	_, err := j.jenkins.CreateJobInFolder(context.Background(), *content, *jobName, *folderName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateJob implements JenkinsService.
func (j *jenkinsUtilImpl) UpdateJob(ctx context.Context, jobName, folderName, content *string) error {
	job, err := j.jenkins.GetJob(ctx, *jobName, *folderName)
	if err != nil {
		return err
	}
	err = job.UpdateConfig(ctx, *content)
	if err != nil {
		return err
	}
	return nil
}

// DeleteJob implements JenkinsService.
func (j *jenkinsUtilImpl) DeleteJob(ctx context.Context, jobName *string, folderName *string) error {
	job, err := j.jenkins.GetJob(ctx, *jobName, *folderName)
	if err != nil {
		return err
	}
	_, err = job.Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
