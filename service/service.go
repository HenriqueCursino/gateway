package service

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/common/env"
	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/dto"
	"github.com/henriquecursino/gateway/middleware"
	"github.com/henriquecursino/gateway/repository"
	"github.com/henriquecursino/gateway/tools"
)

type Service interface {
	UserService(dto.UserRequest) error
	LoginService(loginRequest dto.UserLogin) (*model.Users, error)
	CreateJWT(user *model.Users) (string, error)
	CheckPermission(ctx *gin.Context) bool
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo,
	}
}

func (serv *service) UserService(userRequest dto.UserRequest) error {
	documentUnmasked := tools.RemoveMask(userRequest.Document)

	user := dto.UserCreate{
		FullName: userRequest.FullName,
		Hash:     tools.GenerateHash(),
		Email:    userRequest.Email,
		Document: documentUnmasked,
		Password: userRequest.Password,
		RoleId:   userRequest.RoleID,
	}

	err := serv.repo.CreateUser(&user)
	return err
}

func (serv *service) LoginService(loginRequest dto.UserLogin) (*model.Users, error) {
	login := dto.UserLogin{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	user, err := serv.repo.LoginUser(login)
	if err != nil {
		return nil, err
	}
	if user.Password != login.Password {
		panic("Wrong password!")
	}
	return &user, nil
}

func (serv *service) CreateJWT(user *model.Users) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	oneDay := 1440 // 24 hours

	claims["email"] = user.Email
	claims["userHash"] = user.Hash
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(oneDay)).Unix() // minutes

	tokenString, err := token.SignedString([]byte(env.SecretKeyJWT))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (serv *service) CheckPermission(ctx *gin.Context) bool {
	hash := getHashFromToken(ctx)
	userObj, _ := serv.repo.GetUser(hash)
	permissions, _ := serv.repo.GetAllPermissionsRole(userObj.RoleId)
	for i := 0; i < len(permissions); i++ {
		valid, _ := serv.repo.CheckPermission(permissions[i].ID, common.UserCreate)
		return valid
	}
	return false
}

func getHashFromToken(ctx *gin.Context) string {
	claims, findBody := middleware.DecodedToken(ctx)
	if !findBody {
		ctx.JSON(http.StatusBadGateway, "jwt not found!")
		return ""
	}

	hashInterface := claims[common.KeyHashToken]
	hashString := tools.GetStringFromBody(hashInterface)

	return hashString
}
