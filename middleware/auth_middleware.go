package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/common/env"
	"github.com/henriquecursino/gateway/repository"
	"github.com/henriquecursino/gateway/tools"
)

type Middleware interface {
	CheckPermission(permissionName string) gin.HandlerFunc
}

type middleware struct {
	repo repository.Repository
}

func NewMiddleware(repo repository.Repository) Middleware {
	return &middleware{
		repo,
	}
}

func Validate() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !hasTokenOnHeaders(context) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Token is not in header!")
			return
		}

		token := context.Request.Header.Get(common.HeaderKey)
		if !isValidSignatureToken(token) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Token invalid signature!")
			return
		}

		if !isExpiredToken(context) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Token expired!")
			return
		}

		context.Next()
	}
}

func hasTokenOnHeaders(ctx *gin.Context) bool {
	token := ctx.Request.Header.Get(common.HeaderKey)
	return token != ""
}

func isValidSignatureToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token)
		}
		return []byte(env.SecretKeyJWT), nil
	})

	return err == nil
}

func isExpiredToken(ctx *gin.Context) bool {
	token, _ := decodedToken(ctx)

	exp := token[common.KeyExpToken].(float64)
	test := int64(exp) > time.Now().Unix()
	return test
}

func decodedToken(ctx *gin.Context) (jwt.MapClaims, bool) {
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

func (serv *middleware) CheckPermission(permissionName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hash := GetHashFromToken(ctx)
		userObj, _ := serv.repo.GetUser(hash)
		permissions, _ := serv.repo.GetAllPermissionsRole(userObj.RoleId)
		if len(permissions) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "User doenst have permission!")
			return
		}
		var valid bool
		for i := 0; i < len(permissions); i++ {
			valid, _ = serv.repo.CheckPermissionRepository(permissions[i].ID, permissionName)
			if !valid {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, "User doenst have permission!")
				return
			}
		}
		ctx.Next()
	}
}

func GetHashFromToken(ctx *gin.Context) string {
	claims, findBody := decodedToken(ctx)
	if !findBody {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, "jwt not found!")
		return ""
	}

	hashInterface := claims[common.KeyHashToken]
	hashString := tools.GetStringFromBody(hashInterface)

	return hashString
}
