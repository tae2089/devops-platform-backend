package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

func (s *slackUtil) GenerateModalRequest() slack.ModalViewRequest {
	// Create a ModalViewRequest with a header and two inputs
	titleText := slack.NewTextBlockObject(slack.PlainTextType, "점심 기록하기", false, false)
	closeText := slack.NewTextBlockObject(slack.PlainTextType, "취소", false, false)
	submitText := slack.NewTextBlockObject(slack.PlainTextType, "제출", false, false)

	headerText := slack.NewTextBlockObject("mrkdwn", "오늘도 즐거운 점심이셨나요?", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	//결제한 사람 고르기
	payerSelectText := slack.NewTextBlockObject(slack.PlainTextType, "💳 결제한 사람은 누구인가요?", false, false)
	payerSelectPlaceholder := slack.NewTextBlockObject(slack.PlainTextType, "클릭하여 선택해주세요", false, false)
	payerSelectElement := slack.NewOptionsSelectBlockElement("users_select", payerSelectPlaceholder, "payer_select")
	payerSelect := slack.NewInputBlock("payer selector", payerSelectText, nil, payerSelectElement)

	//식사를 같이한 사람들
	userSelectText := slack.NewTextBlockObject(slack.PlainTextType, "👩🏻‍🤝‍👩🏻 같이 먹은 사람을 입력해주세요", false, false)
	userSelectPlaceholder := slack.NewTextBlockObject(slack.PlainTextType, "여러명인 경우 모두 선택해주세요", false, false)
	userSelectElement := slack.NewOptionsMultiSelectBlockElement("multi_users_select", userSelectPlaceholder, "users_select")
	userSelect := slack.NewInputBlock("users selector", userSelectText, nil, userSelectElement)

	// 식사한 장소
	restaurantNameText := slack.NewTextBlockObject(slack.PlainTextType, "🍕 식당 이름을 입력해주세요*", false, false)
	restaurantPlaceholder := slack.NewTextBlockObject(slack.PlainTextType, "TIP -  '배민' 이용 시, '배달의 민족'을 입력해주세요. ", false, false)
	restaurantNameElement := slack.NewPlainTextInputBlockElement(restaurantPlaceholder, "restaurantName")
	restaurantName := slack.NewInputBlock("Restaurant Name", restaurantNameText, nil, restaurantNameElement)

	// 카페를 먹었을떄 사용
	cafeNameText := slack.NewTextBlockObject(slack.PlainTextType, "☕ 카페에 가셨다면, 카페 이름도 입력해주세요", false, false)
	cafeNamePlaceholder := slack.NewTextBlockObject(slack.PlainTextType, "입력해주세요", false, false)
	cafeNameElement := slack.NewPlainTextInputBlockElement(cafeNamePlaceholder, "cafeName")
	cafeName := slack.NewInputBlock("Cafe Name", cafeNameText, nil, cafeNameElement)
	cafeName.Optional = true

	// 결제일자
	paymentDateText := slack.NewTextBlockObject(slack.PlainTextType, "📅결제일을 입력해주세요", false, false)
	paymentDatePlaceholder := slack.NewTextBlockObject(slack.PlainTextType, "숫자 8자리로 입력해주세요. ex) 20230108", false, false)
	paymentDateElement := slack.NewPlainTextInputBlockElement(paymentDatePlaceholder, "paymentDate")
	paymentDateName := slack.NewInputBlock("Payment Date", paymentDateText, nil, paymentDateElement)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			payerSelect,
			userSelect,
			restaurantName,
			cafeName,
			paymentDateName,
		},
	}

	var modalRequest slack.ModalViewRequest
	modalRequest.Type = "modal"
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks
	modalRequest.PrivateMetadata = "/slash"
	return modalRequest
}

func (s *slackUtil) GenerateFrontDeployModal(projects ...string) slack.ModalViewRequest {
	// Create a ModalViewRequest with a header and two inputs
	titleText := slack.NewTextBlockObject(slack.PlainTextType, "Front 배포", false, false)
	closeText := slack.NewTextBlockObject(slack.PlainTextType, "취소", false, false)
	submitText := slack.NewTextBlockObject(slack.PlainTextType, "제출", false, false)
	headerText := slack.NewTextBlockObject(slack.MarkdownType, "Front Web 배포하기", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	options := s.getExternalSelctOption(projects, false)

	projectSelector := s.getSelectBlock("배포할 프로젝트를 선택해주세요?", "클릭하여 선택해주세요", "project_select", "project selector", false, options)
	//github repository 입력하기
	gitRepoPlainTextBlock := s.getPlainTextBlock("배포에 사용될 레포지토리를 입력해주세요.", "레포지토링명을 입력해주세요. ex) main, develop", "repository", "repository block")
	// // 브랜치명 입력하기
	branchPlainTextBlock := s.getPlainTextBlock("배포에 사용될 브랜치를 입력해주세요.", "브랜치명을 입력해주세요. ex) main, develop", "branch", "branch block")
	// // 배포 시 사용할 도메인 입력하기
	domainPlainTextBlock := s.getPlainTextBlock("배포된 웹사이트에 사용할 도메인을 입력해주세요.", "사용하실 도메인을 입력해주세요. ex)www.example.com", "domain", "domain block")
	// // certificate arn 입력하기
	cetificateArnBlock := s.getPlainTextBlock("aws certificate arn을 입력해주세요.", "arn을 입력해주세요", "certificateArn", "certificate block")

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			projectSelector,
			gitRepoPlainTextBlock,
			branchPlainTextBlock,
			domainPlainTextBlock,
			cetificateArnBlock,
		},
	}
	metadata := "/front-deploy"
	modalRequest := s.createModalRequest(titleText, closeText, submitText, blocks, metadata)
	return modalRequest
}

func (s *slackUtil) createModalRequest(titleText *slack.TextBlockObject, closeText *slack.TextBlockObject, submitText *slack.TextBlockObject, blocks slack.Blocks, metadata string) slack.ModalViewRequest {
	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.VTModal
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks
	modalRequest.PrivateMetadata = metadata
	return modalRequest
}

func (s *slackUtil) getPlainTextBlock(text, placeholder, actionID, blockID string) *slack.InputBlock {
	blockText := slack.NewTextBlockObject(slack.PlainTextType, text, false, false)
	blockPlaceholder := slack.NewTextBlockObject(slack.PlainTextType, placeholder, false, false)
	blockElement := slack.NewPlainTextInputBlockElement(blockPlaceholder, actionID)
	blockName := slack.NewInputBlock(blockID, blockText, nil, blockElement)
	return blockName
}

func (s *slackUtil) getMarkdownBlock(text, placeholder, actionID, blockID string) *slack.InputBlock {
	blockText := slack.NewTextBlockObject(slack.MarkdownType, text, false, false)
	blockPlaceholder := slack.NewTextBlockObject(slack.MarkdownType, placeholder, false, false)
	blockElement := slack.NewPlainTextInputBlockElement(blockPlaceholder, actionID)
	blockName := slack.NewInputBlock(blockID, blockText, nil, blockElement)
	return blockName
}

func (s *slackUtil) getSelectBlock(text, placeholder, actionID, blockID string, optional bool, options []*slack.OptionBlockObject) *slack.InputBlock {
	selectText := slack.NewTextBlockObject(slack.PlainTextType, text, false, false)
	selectPlaceholder := slack.NewTextBlockObject(slack.PlainTextType, placeholder, false, false)
	selectElement := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, nil, actionID, options...)
	selector := slack.NewInputBlock(blockID, selectText, selectPlaceholder, selectElement)

	// inviteeOption := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, nil, "invitee", options...)
	// inviteeBlock := slack.NewInputBlock("invitee", inviteeText, nil, inviteeOption)

	// inviteeText := slack.NewTextBlockObject(slack.PlainTextType, "Invitee from static list", false, false)
	// inviteeOption := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, nil, "invitee", option_test...)
	// inviteeBlock := slack.NewInputBlock("invitee", inviteeText, nil, inviteeOption)
	return selector
}

func (s *slackUtil) getMultiSelectUser(text, placeholder, actionID, blockID string) *slack.InputBlock {
	userSelectText := slack.NewTextBlockObject(slack.PlainTextType, text, false, false)
	userSelectPlaceholder := slack.NewTextBlockObject(slack.PlainTextType, placeholder, false, false)
	userSelectElement := slack.NewOptionsMultiSelectBlockElement("multi_users_select", userSelectPlaceholder, actionID)
	userSelect := slack.NewInputBlock(blockID, userSelectText, nil, userSelectElement)
	return userSelect
}

// createOptionBlockObjects - utility function for generating option block objects
func (s *slackUtil) getExternalSelctOption(options []string, users bool) []*slack.OptionBlockObject {
	optionBlockObjects := make([]*slack.OptionBlockObject, 0, len(options))
	var text string
	for _, o := range options {
		if users {
			text = fmt.Sprintf("<@%s>", o)
		} else {
			text = o
		}
		optionText := slack.NewTextBlockObject(slack.PlainTextType, text, false, false)
		optionBlockObjects = append(optionBlockObjects, slack.NewOptionBlockObject(o, optionText, nil))
	}
	return optionBlockObjects
}
