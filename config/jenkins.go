package config

type Jenkins struct {
	URL      string `env:"JENKINS_URL,required"`
	User     string `env:"JENKINS_USER,required"`
	PassWord string `env:"JENKINS_PASSWORD,required"`
}
