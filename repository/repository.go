package repository

import (
	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/dto"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *dto.UserCreate) error
	LoginUser(login dto.UserLogin) (model.Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (repo *repository) CreateUser(user *dto.UserCreate) error {
	if err := repo.db.Table(model.TableUserName).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) LoginUser(login dto.UserLogin) (model.Users, error) {
	user := model.Users{}
	err := repo.db.Table(model.TableUserName).Where("email = ?", login.Email).First(&user).Error
	return user, err
}
