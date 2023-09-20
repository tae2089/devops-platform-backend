package config

import (
	"github.com/tae2089/devops-platform-backend/domain"
	"os"
)

var (
	jwtKey domain.JwtKey
)

func GetJwtKey() domain.JwtKey {
	if jwtKey.PrivateKey == "" || jwtKey.PublicKey == "" {
		jwtKey.PrivateKey = os.Getenv("PRIVATE_KEY")
		jwtKey.PublicKey = os.Getenv("PUBLIC_KEY")
	}
	return jwtKey
}
