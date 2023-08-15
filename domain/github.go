package domain

type RequestGithubWebhookDto struct {
	Owner     string
	Repo      string
	Token     string
	TargetUrl string
}
