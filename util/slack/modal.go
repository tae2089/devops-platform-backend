package slack

import "github.com/slack-go/slack"

func (s *slackUtilImpl) GenerateModalRequest() slack.ModalViewRequest {
	// Create a ModalViewRequest with a header and two inputs
	titleText := slack.NewTextBlockObject("plain_text", "점심 기록하기", false, false)
	closeText := slack.NewTextBlockObject("plain_text", "취소", false, false)
	submitText := slack.NewTextBlockObject("plain_text", "제출", false, false)

	headerText := slack.NewTextBlockObject("mrkdwn", "오늘도 즐거운 점심이셨나요?", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	//결제한 사람 고르기
	payerSelectText := slack.NewTextBlockObject("plain_text", "💳 결제한 사람은 누구인가요?", false, false)
	payerSelectPlaceholder := slack.NewTextBlockObject("plain_text", "클릭하여 선택해주세요", false, false)
	payerSelectElement := slack.NewOptionsSelectBlockElement("users_select", payerSelectPlaceholder, "payer_select")
	payerSelect := slack.NewInputBlock("payer selector", payerSelectText, nil, payerSelectElement)

	//식사를 같이한 사람들
	userSelectText := slack.NewTextBlockObject("plain_text", "👩🏻‍🤝‍👩🏻 같이 먹은 사람을 입력해주세요", false, false)
	userSelectPlaceholder := slack.NewTextBlockObject("plain_text", "여러명인 경우 모두 선택해주세요", false, false)
	userSelectElement := slack.NewOptionsMultiSelectBlockElement("multi_users_select", userSelectPlaceholder, "users_select")
	userSelect := slack.NewInputBlock("users selector", userSelectText, nil, userSelectElement)

	// 식사한 장소
	restaurantNameText := slack.NewTextBlockObject("plain_text", "🍕 식당 이름을 입력해주세요*", false, false)
	restaurantPlaceholder := slack.NewTextBlockObject("plain_text", "TIP -  '배민' 이용 시, '배달의 민족'을 입력해주세요. ", false, false)
	restaurantNameElement := slack.NewPlainTextInputBlockElement(restaurantPlaceholder, "restaurantName")
	restaurantName := slack.NewInputBlock("Restaurant Name", restaurantNameText, nil, restaurantNameElement)

	// 카페를 먹었을떄 사용
	cafeNameText := slack.NewTextBlockObject("plain_text", "☕ 카페에 가셨다면, 카페 이름도 입력해주세요", false, false)
	cafeNamePlaceholder := slack.NewTextBlockObject("plain_text", "입력해주세요", false, false)
	cafeNameElement := slack.NewPlainTextInputBlockElement(cafeNamePlaceholder, "cafeName")
	cafeName := slack.NewInputBlock("Cafe Name", cafeNameText, nil, cafeNameElement)
	cafeName.Optional = true

	// 결제일자
	paymentDateText := slack.NewTextBlockObject("plain_text", "📅결제일을 입력해주세요", false, false)
	paymentDatePlaceholder := slack.NewTextBlockObject("plain_text", "숫자 8자리로 입력해주세요. ex) 20230108", false, false)
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
