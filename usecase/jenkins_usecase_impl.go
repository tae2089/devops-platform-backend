package usecase

import (
	"context"
	"net/http"

	"github.com/tae2089/devops-platform-backend/exception"
	"github.com/tae2089/devops-platform-backend/repository"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
)

var _ (JenkinsUsecase) = (*jenkinsUsecaseImpl)(nil)

type jenkinsUsecaseImpl struct {
	jenkinsUtil       jenkins.Util
	githubUtil        github.Util
	userRepository    repository.UserRepository
	jenkinsRepository repository.JenkinsRepository
}

// RegisteProject implements JenkinsUsecase.
func (j *jenkinsUsecaseImpl) RegistProject(request *http.Request) error {

	user, err := j.userRepository.FindBySlackID(context.Background(), "")

	if err != nil {
		return exception.IsEntityNotFound(err)
	}
	if err = exception.IsAdminRole(user.Roles); err != nil {
		return err
	}
	return nil
}

// RegistFrontJob implements JenkinsUsecase.
func (j *jenkinsUsecaseImpl) RegistJob(request *http.Request) error {
	// ctx := context.Background()

	return nil
}
