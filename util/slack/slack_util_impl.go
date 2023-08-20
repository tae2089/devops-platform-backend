package slack

import (
	"fmt"
	"log"
	"net/http"

	"github.com/slack-go/slack"
)

type slackUtilImpl struct {
	client *slack.Client
}

var _ SlackUtil = (*slackUtilImpl)(nil)

func (s *slackUtilImpl) OpenView(triggerId string, modalRequest slack.ModalViewRequest) error {
	_, err := s.client.OpenView(triggerId, modalRequest)
	if err != nil {
		log.Printf("Error opening view: %s", err)
		return err
	}
	return nil
}

func (s *slackUtilImpl) PostMessage(channelId string, options ...slack.MsgOption) error {
	_, _, err := s.client.PostMessage(channelId, options...)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}

func (s *slackUtilImpl) GetUserProfile(userId string) (string, error) {
	userProfile, err := s.client.GetUserProfile(
		&slack.GetUserProfileParameters{
			UserID: userId,
		},
	)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(userProfile.RealName)
	log.Println(userProfile.DisplayName)
	return userProfile.RealName, nil
}

func (s *slackUtilImpl) GetUsersRealName(userId ...string) ([]string, error) {
	users, err := s.client.GetUsersInfo(userId...)
	userRealNameList := []string{}
	if err != nil {
		return userRealNameList, err
	}
	for _, user := range *users {
		userRealNameList = append(userRealNameList, user.RealName)
	}
	return userRealNameList, nil
}

func (s *slackUtilImpl) SlashCommandParse(request *http.Request) (string, error) {

	slackCommand, err := slack.SlashCommandParse(request)
	if err != nil {
		return "", err
	}
	log.Println("userID", slackCommand.UserID)
	return slackCommand.TriggerID, nil
}
