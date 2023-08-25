package usecase

import (
	"github.com/slack-go/slack"
	"github.com/tae2089/devops-platform-backend/domain"
	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
	slackUtil "github.com/tae2089/devops-platform-backend/util/slack"
)

var _ (SlackUsecase) = (*slackUsecase)(nil)

type slackUsecase struct {
	slackUtil   slackUtil.Util
	githubUtil  github.Util
	jenkinsUtil jenkins.Util
}

// RegistJenkinsFrontJob implements SlackUsecase.
func (s *slackUsecase) RegistJenkinsFrontJob(callback *slack.InteractionCallback) error {
	projectSelect := callback.View.State.Values["project selector"]["project_select"]
	repositoryField := callback.View.State.Values["repository block"]["repository"]
	branchField := callback.View.State.Values["branch block"]["branch"]
	webhookField := callback.View.State.Values["webhook block"]["webhook"]
	jenkinsFileField := callback.View.State.Values["jenkinsfile block"]["jenkinsfile"]
	jenkinsJobField := callback.View.State.Values["jenkinsjob block"]["jenkinsjob"]
	content := s.jenkinsUtil.GetJenkinsJobContent(repositoryField.Value, branchField.Value, projectSelect.SelectedOption.Value, webhookField.Value, jenkinsFileField.Value)
	_, err := s.jenkinsUtil.CreateJob(&jenkinsJobField.Value, &projectSelect.SelectedOption.Text.Text, &content)
	if err != nil {
		return err
	}
	responseJob := &domain.ResultMessageJenkinsJob{
		Project:      projectSelect.SelectedOption.Value,
		WebhookToken: webhookField.Value,
		GitURL:       repositoryField.Value,
		Branch:       branchField.Value,
		FileName:     jenkinsFileField.Value,
	}
	channelID := callback.Channel.ID
	if callback.Channel.Name == "" {
		channelID = callback.User.ID
	}

	resultBlocks := s.slackUtil.GetJenkinsJobResultBlocks(responseJob.Write())
	err = s.slackUtil.PostMessageWithBlocks(channelID, resultBlocks)
	if err != nil {
		return err
	}
	return nil
}

// GetCallbackPayload implements SlackUsecase.
func (s *slackUsecase) GetCallbackPayload(payload *string) (*slack.InteractionCallback, error) {
	i, err := s.slackUtil.GetCallbackPayload(payload)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// func (j *jenkinsUsecaseImpl) SaveLunchPayment(i slack.InteractionCallback, token string) error {

// 	payerSelect := i.View.State.Values["payer selector"]["payer_select"]
// 	userSelect := i.View.State.Values["users selector"]["users_select"]
// 	restaurantField := i.View.State.Values["Restaurant Name"]["restaurantName"]
// 	cafeField := i.View.State.Values["Cafe Name"]["cafeName"]
// 	paymentDateField := i.View.State.Values["Payment Date"]["paymentDate"]
// 	paymentDate, err := time.Parse("20060102", paymentDateField.Value)
// 	if err != nil {
// 		return err
// 	}
// 	payerName, err := local.GetUserProfile(payerSelect.SelectedUser)
// 	if err != nil {
// 		return err
// 	}
// 	users, err := local.GetUsersRealName(userSelect.SelectedUsers...)
// 	if err != nil {
// 		return err
// 	}
// 	log.Println("users", users)
// 	participant, err := c.participantRepository.SaveParticipants(users)
// 	if err != nil {
// 		return err
// 	}

// 	lunchDto := dto.NewLunchDto(payerName, restaurantField.Value, cafeField.Value, paymentDate)

// 	_ = c.lunchRepository.SaveLunchPayment(*lunchDto, participant)
// 	msg := fmt.Sprintf("%s 님 점심 결제 등록이 정상적으로 완료됐습니다.!", payerName)

// 	err = local.PostMessage(payerSelect.SelectedUser,
// 		slack.MsgOptionText(msg, false),
// 		slack.MsgOptionAttachments(),
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
