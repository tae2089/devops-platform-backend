package jenkins

import (
	"context"

	"github.com/bndr/gojenkins"
	"github.com/tae2089/devops-platform-backend/config"
)

type JenkinsService interface {
	CreateJob(jobName, folderName, content *string) (*gojenkins.Job, error)
	UpdateJob(ctx context.Context, jobName, folderName, content *string) (*gojenkins.Job, error)
	DeleteJob(ctx context.Context, jobName, folderName *string) error
}

func NewJenkinsService(jenkinsCOnfig *config.Jenkins) JenkinsService {
	jenkinsConfig := config.GetJenkinsConfig()
	jenkinsClient := gojenkins.CreateJenkins(nil, jenkinsConfig.URL, jenkinsConfig.User, jenkinsConfig.PassWord)
	_, err := jenkinsClient.Init(context.Background())
	if err != nil {
		panic(err)
	}
	return &jenkinsServiceImpl{
		jenkins: jenkinsClient,
	}
}
