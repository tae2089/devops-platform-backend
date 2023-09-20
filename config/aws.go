package config

type Aws struct {
	Profiles []string `env:"PROFILES" envSeparator:";"`
}
