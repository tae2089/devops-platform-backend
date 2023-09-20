package domain

import "github.com/golang-jwt/jwt/v5"

type JwtKey struct {
	PrivateKey string
	PublicKey  string
}

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}
