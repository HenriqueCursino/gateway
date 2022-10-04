package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/henriquecursino/gateway/common/env"
	"github.com/henriquecursino/gateway/dto"
	"github.com/henriquecursino/gateway/repository"
	"github.com/henriquecursino/gateway/tools"
)

type Service interface {
	UserService(dto.UserRequest) error
	LoginService(loginRequest dto.UserLogin) error
	CreateJWT(user *dto.UserLogin) (string, error)
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) Service {
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

func (serv *service) LoginService(loginRequest dto.UserLogin) error {
	login := dto.UserLogin{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	user, err := serv.repo.LoginUser(login)
	if err != nil {
		return err
	}
	if user.Password != login.Password {
		panic("Wrong password!")
	}
	return nil
}

func (serv *service) CreateJWT(user *dto.UserLogin) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	oneDay := 1440 // 24 hours

	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(oneDay)).Unix() // minutes

	tokenString, err := token.SignedString([]byte(env.SecretKeyJWT))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
