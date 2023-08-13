package config

type Github struct {
	Token string `env:"GITHUB_TOKEN,required"`
}
