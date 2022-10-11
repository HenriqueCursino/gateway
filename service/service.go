package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/common/env"
	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/dto"
	"github.com/henriquecursino/gateway/repository"
	"github.com/henriquecursino/gateway/tools"
)

type Service interface {
	UserService(dto.UserRequest) error
	LoginService(loginRequest dto.UserLogin) (*model.Users, error)
	CreateJWT(user *model.Users) (string, error)
	GetAllUsersService() ([]dto.AllUsers, error)
	UpdateUserRole(updateUser dto.UpdateUserRole) error
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
		UserId:   tools.GenerateHash(),
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

	oneDay := common.RemaningHoursToExpiredToken

	claims["email"] = user.Email
	claims["userHash"] = user.UserId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(oneDay)).Unix() // minutes

	tokenString, err := token.SignedString([]byte(env.SecretKeyJWT))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (serv *service) GetAllUsersService() ([]dto.AllUsers, error) {
	allUsers, err := serv.repo.GetAllUsers()
	usersReturn := []dto.AllUsers{}
	for i := 0; i < len(allUsers); i++ {
		role, _ := serv.repo.GetRole(allUsers[i].RoleId)

		usersReturn = append(usersReturn, dto.AllUsers{
			FullName: allUsers[i].FullName,
			UserId:   allUsers[i].UserId,
			Email:    allUsers[i].Email,
			Document: allUsers[i].Document,
			Password: allUsers[i].Password,
			Roles:    role,
		})
	}

	if err != nil {
		return []dto.AllUsers{}, err
	}

	return usersReturn, nil
}

func (serv *service) UpdateUserRole(updateUser dto.UpdateUserRole) error {
	err := serv.repo.UpdateUserRole(updateUser.Document, updateUser.NewRole)
	return err
}
