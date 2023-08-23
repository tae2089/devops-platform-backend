package jenkins

import (
	"context"
	"fmt"

	"github.com/tae2089/bob-logging/logger"

	"github.com/bndr/gojenkins"
	"github.com/tae2089/devops-platform-backend/domain"
)

var _ Util = (*utilImpl)(nil)

type utilImpl struct {
	jenkins *gojenkins.Jenkins
}

// GetJenkinsJobContent implements Util.
func (j *utilImpl) GetJenkinsJobContent(repository, branch, project, webhook, jenkinsFile string) string {

	url := fmt.Sprintf("git@github.com:%s.git", repository)
	jenkinsJob := domain.JenkinsJob{
		GitURL:       url,
		BranchName:   branch,
		TeamName:     project,
		FileName:     jenkinsFile,
		WebhookToken: webhook,
	}
	return jenkinsJob.Write()
}

// GetFrontJobContent implements Util.
func (j *utilImpl) GetJenkinsFrontFileContent(repository string, branch string, slackChannel string, buckName string, CloudFrontID string, project string) string {
	jenkinsJob := &domain.JenkinsFrontFile{
		GitURL:       fmt.Sprintf("git@github.com:%s.git", repository),
		Branch:       branch,
		SlackChannel: slackChannel,
		BuckName:     buckName,
		CloudFrontID: CloudFrontID,
		AwsProfile:   project,
	}
	return jenkinsJob.Write()
}

func (j *utilImpl) CreateJob(jobName, folderName, content *string) (*gojenkins.Job, error) {
	_, err := j.jenkins.CreateJobInFolder(context.Background(), *content, *jobName, *folderName)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return nil, nil
}

// UpdateJob implements JenkinsService.
func (j *utilImpl) UpdateJob(ctx context.Context, jobName, folderName, content *string) error {
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
func (j *utilImpl) DeleteJob(ctx context.Context, jobName *string, folderName *string) error {
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
