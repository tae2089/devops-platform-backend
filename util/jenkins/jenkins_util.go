package jenkins

import (
	"context"
	"log"

	"github.com/bndr/gojenkins"
	"github.com/tae2089/devops-platform-backend/config"
)

type Util interface {
	CreateJob(jobName, folderName, content *string) (*gojenkins.Job, error)
	UpdateJob(ctx context.Context, jobName, folderName, content *string) error
	DeleteJob(ctx context.Context, jobName, folderName *string) error
	GetJenkinsFrontFileContent(gitURL, branch, slackChannel, buckName, CloudFrontID, project string) string
	GetJenkinsJobContent(repository, branch, project, webhook, jenkinsFile string) string
}

func NewJenkinsUtil(jenkinsConfig *config.Jenkins) Util {
	jenkinsClient := gojenkins.CreateJenkins(nil, jenkinsConfig.URL, jenkinsConfig.User, jenkinsConfig.PassWord)
	_, err := jenkinsClient.Init(context.Background())
	if err != nil {
		log.Println(err)
		return nil
	}
	return &utilImpl{
		jenkins: jenkinsClient,
	}
}
