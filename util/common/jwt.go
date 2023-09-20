package common

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tae2089/devops-platform-backend/config"
	"github.com/tae2089/devops-platform-backend/domain"
	"log"
	"time"
)

func CreateAccessToken(payload []byte) (accessToken string, err error) {
	jwtKey := config.GetJwtKey()
	privateKeyBytes := []byte(jwtKey.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 개인 키 파싱
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("Failed to decode PEM block containing RSA private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	claims := domain.JwtCustomClaims{
		Name: "bc-labs-token",
		ID:   uuid.NewString(),
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365 * 10)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "devops-team",
			Subject:   bytes.NewBuffer(payload).String(),
			Audience:  []string{"bc-labs-project"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	if err != nil {
		return "", err
	}
	t, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return t, err
}

func IsAuthorized(requestToken string) (bool, error) {
	_, err := parseToken(requestToken)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractIDFromToken(requestToken string) (string, error) {
	token, err := parseToken(requestToken)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["id"].(string), nil
}

func parseToken(requestToken string) (*jwt.Token, error) {
	jwtKey := config.GetJwtKey()
	publicKeyBytes := []byte(jwtKey.PublicKey)
	block, _ := pem.Decode(publicKeyBytes)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}
