package config

type SlackBot struct {
	AccessToken string `env:"SLACK_BOT_ACCESS_TOKEN,required"`
	SecretToken string `env:"SLACK_BOT_SECRET_TOKEN,required"`
}
