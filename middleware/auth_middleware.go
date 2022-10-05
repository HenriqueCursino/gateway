package middleware

import (
	"fmt"
	"log"
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

func DecodedToken(ctx *gin.Context) (jwt.MapClaims, bool) {
	stringToken := ctx.GetHeader(common.HeaderKey)
	secrectKey := env.SecretKeyJWT
	hmacSecret := []byte(secrectKey)

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
