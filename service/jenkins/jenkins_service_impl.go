package jenkins

import (
	"context"

	"github.com/bndr/gojenkins"
)

type jenkinsServiceImpl struct {
	jenkins *gojenkins.Jenkins
}

var _ JenkinsService = (*jenkinsServiceImpl)(nil)

func (j *jenkinsServiceImpl) CreateJob(jobName, folderName, content *string) (*gojenkins.Job, error) {
	_, err := j.jenkins.CreateJobInFolder(context.Background(), *content, *jobName, *folderName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateJob implements JenkinsService.
func (j *jenkinsServiceImpl) UpdateJob(ctx context.Context, jobName, folderName, content *string) (*gojenkins.Job, error) {
	jobPath := *folderName + "/" + *jobName
	_ = j.jenkins.UpdateJob(ctx, jobPath, *content)
	return nil, nil
}
