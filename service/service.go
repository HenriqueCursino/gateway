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
	DeleteUserService(user dto.UserDelete) error
	CreateRole(newRole dto.RoleUser) error
	GetAllRoles() ([]dto.AllRoles, error)
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
	passwordHash, _ := tools.Encrypt(userRequest.Password)

	user := dto.UserCreate{
		FullName: userRequest.FullName,
		UserId:   tools.GenerateHash(),
		Email:    userRequest.Email,
		Document: documentUnmasked,
		Password: *passwordHash,
		RoleId:   userRequest.RoleID,
	}

	err := serv.repo.CreateUser(&user)
	return err
}

func (serv *service) LoginService(loginRequest dto.UserLogin) (*model.Users, error) {
	passwordHash, _ := tools.Encrypt(loginRequest.Password)

	login := dto.UserLogin{
		Email:    loginRequest.Email,
		Password: *passwordHash,
	}

	user, err := serv.repo.LoginUser(login)
	if err != nil {
		return nil, err
	}

	if tools.Compare(user.Password, loginRequest.Password) {
		return &user, nil
	}
	panic("Wrong password!")
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

func (serv *service) DeleteUserService(user dto.UserDelete) error {
	if err := serv.repo.DeleteUser(user.UserId); err != nil {
		return err
	}
	return nil
}

func (serv *service) CreateRole(newRole dto.RoleUser) error {
	if err := serv.repo.CreateRole(newRole); err != nil {
		return err
	}
	return nil
}

func (serv *service) GetAllRoles() ([]dto.AllRoles, error) {
	allRoles := serv.repo.GetAllRoles()
	allRolesResponse := []dto.AllRoles{}
	permissions := []dto.Permissions{}

	for i := 0; i < len(allRoles); i++ {
		allPermRole, err := serv.repo.GetAllPermissionsRole(allRoles[i].ID)
		if err != nil {
			return allRolesResponse, err
		}
		for index := 0; index < len(allPermRole); index++ {
			allPerm := serv.repo.GetAllPermissions(allPermRole[index].PermissionId)
			permissions = append(permissions, dto.Permissions{
				Permission: allPerm.Permission,
			})
		}
		allRolesResponse = append(allRolesResponse, dto.AllRoles{
			Role:        allRoles[i].Role,
			Permissions: permissions,
		})
		permissions = []dto.Permissions{}
	}
	return allRolesResponse, nil
}
