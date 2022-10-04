package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/common/env"
)

func Validate() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !hasTokenOnHeaders(context) {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := context.Request.Header.Get(common.HeaderKey)
		if !isValidToken(token) {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		context.Next()
	}
}

func hasTokenOnHeaders(ctx *gin.Context) bool {
	token := ctx.Request.Header.Get(common.HeaderKey)
	return token != ""
}

func isValidToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token)
		}
		return []byte(env.SecretKeyJWT), nil
	})

	return err == nil
}
