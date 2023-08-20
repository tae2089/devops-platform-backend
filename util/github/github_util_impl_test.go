package github_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/domain"
	"github.com/tae2089/devops-platform-backend/util/github"
)

func Test_gitServiceImpl_UploadFile(t *testing.T) {
	gitConfig := config.GetGithubConfig()
	g := github.NewGithubUtil(gitConfig)
	t.Run("upload file", func(t *testing.T) {
		t.Skip("")
		hookDto := &domain.RequestUploadFileDto{
			Owner:   "<GITUHB USER ID>",
			Repo:    "<GITUH REPON NAME>",
			Path:    "<FILE PATH>",
			Branch:  "<BRANCH>",
			Content: []byte("<CONTENT>"),
		}
		err := g.UploadFile(context.Background(), hookDto)
		assert.Nil(t, err)
	})

	t.Run("not found file", func(t *testing.T) {
		t.Skip("")
		hookDto := &domain.RequestUploadFileDto{
			Owner:   "<GITUHB USER ID>",
			Repo:    "<GITUH REPON NAME>",
			Path:    "<FILE PATH>",
			Branch:  "<BRANCH>",
			Content: []byte("<CONTENT>"),
		}
		err := g.UploadFile(context.Background(), hookDto)
		log.Println(err)
		assert.Nil(t, err)
	})
}
