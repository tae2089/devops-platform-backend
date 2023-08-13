package domain

type RegisterGithubWebhookDto struct {
	Owner     string
	Repo      string
	Token     string
	TargetUrl string
}
