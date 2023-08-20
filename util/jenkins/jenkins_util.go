package jenkins

import (
	"context"

	"github.com/bndr/gojenkins"
	"github.com/tae2089/devops-platform-backend/config"
)

type JenkinsUtil interface {
	CreateJob(jobName, folderName, content *string) (*gojenkins.Job, error)
	UpdateJob(ctx context.Context, jobName, folderName, content *string) error
	DeleteJob(ctx context.Context, jobName, folderName *string) error
}

func NewJenkinsUtil(jenkinsConfig *config.Jenkins) JenkinsUtil {
	jenkinsClient := gojenkins.CreateJenkins(nil, jenkinsConfig.URL, jenkinsConfig.User, jenkinsConfig.PassWord)
	_, err := jenkinsClient.Init(context.Background())
	if err != nil {
		panic(err)
	}
	return &jenkinsUtilImpl{
		jenkins: jenkinsClient,
	}
}
