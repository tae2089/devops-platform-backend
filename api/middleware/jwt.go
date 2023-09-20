package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/devops-platform-backend/exception"
	"github.com/tae2089/devops-platform-backend/util/common"
	"net/http"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := common.IsAuthorized(authToken)
			if authorized {
				userID, err := common.ExtractIDFromToken(authToken)
				if err != nil {
					c.JSON(http.StatusUnauthorized, exception.NotAuthorized)
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			if err != nil {
				logger.Error(err)
				c.JSON(http.StatusUnauthorized, exception.NotAuthorized)
				c.Abort()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, exception.NotAuthorized)
		c.Abort()
	}
}
