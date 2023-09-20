package config

type DB struct {
	Port     string `env:"DB_PORT,required"`
	Host     string `env:"DB_HOST,required"`
	DBname   string `env:"DB_NAME,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
}
