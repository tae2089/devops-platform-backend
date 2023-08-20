package middleware

import (
	"bytes"
	"io"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"github.com/tae2089/devops-platform-backend/config"
)

func VerifySlack() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := verifySigningSecret(c.Request)
		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func verifySigningSecret(r *http.Request) error {
	verifier, err := slack.NewSecretsVerifier(r.Header, config.GetSlackBotConfig().SecretToken)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// Need to use r.Body again when unmarshalling SlashCommand and InteractionCallback
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	verifier.Write(body)
	if err = verifier.Ensure(); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
