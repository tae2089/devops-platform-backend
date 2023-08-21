package usecase

type SlackUsecase interface{}

func NewSlackUsecase() SlackUsecase {
	return &slackUsecase{}
}
