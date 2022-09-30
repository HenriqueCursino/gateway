package service

import (
	"github.com/henriquecursino/gateway/dto"
	"github.com/henriquecursino/gateway/repository"
	"github.com/henriquecursino/gateway/tools"
)

type Service interface {
	UserService(dto.UserRequest) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo,
	}
}

func (serv service) UserService(userRequest dto.UserRequest) error {
	documentUnmasked := tools.RemoveMask(userRequest.Document)

	user := dto.UserCreate{
		FullName: userRequest.FullName,
		Email:    userRequest.Email,
		Document: documentUnmasked,
		Password: userRequest.Password,
		RoleId:   userRequest.RoleID,
	}

	err := serv.repo.CreateUser(&user)
	return err
}
