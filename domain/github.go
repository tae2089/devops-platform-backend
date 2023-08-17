package domain

type RequestGithubWebhookDto struct {
	Owner     string
	Repo      string
	Token     string
	TargetUrl string
}

type RequestUploadFileDto struct {
	Owner   string
	Repo    string
	Path    string
	Branch  string
	Content []byte
}
