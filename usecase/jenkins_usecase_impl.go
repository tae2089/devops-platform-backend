package usecase

import (
	"fmt"
	"net/http"

	"github.com/tae2089/devops-platform-backend/util/github"
	"github.com/tae2089/devops-platform-backend/util/jenkins"
	"github.com/tae2089/devops-platform-backend/util/slack"
)

var _ (JenkinsUsecase) = (*jenkinsUsecaseImpl)(nil)

type jenkinsUsecaseImpl struct {
	slackUtil   slack.SlackUtil
	jenkinsUtil jenkins.JenkinsUtil
	githubUtil  github.GithubUtil
}

// RegisterLunchPayment implements JenkinsUsecase.
func (j *jenkinsUsecaseImpl) RegisterLunchPayment(request *http.Request) error {
	triggerId, err := j.slackUtil.SlashCommandParse(request)
	if err != nil {
		return err
	}
	modalRequest := j.slackUtil.GenerateModalRequest()
	err = j.slackUtil.OpenView(triggerId, modalRequest)
	if err != nil {
		fmt.Printf("Error opening view: %s", err)
		return err
	}
	return nil
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
