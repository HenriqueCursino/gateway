package repository

import (
	"github.com/henriquecursino/gateway/dto"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *dto.UserCreat) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (repo repository) CreateUser(user *dto.UserCreat) error {
	if err := repo.db.Table("users").Create(&user).Error; err != nil {
		return err
	}
	return nil
}
