package repository

import (
	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/dto"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *dto.UserCreate) error
	LoginUser(login dto.UserLogin) (model.Users, error)
	GetUser(hash string) (model.Users, error)
	GetAllPermissionsRole(roleId int) ([]model.PermissionsRoles, error)
	CheckPermission(permissionId int, namePermission string) (bool, error)
	GetAllUsers() ([]model.Users, error)
	GetRole(roleId int) (model.Roles, error)
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

func (repo *repository) GetUser(hash string) (model.Users, error) {
	user := model.Users{}
	err := repo.db.Table(model.TableUserName).Where("user_id = ?", hash).First(&user).Error
	return user, err
}

func (repo *repository) GetAllPermissionsRole(roleId int) ([]model.PermissionsRoles, error) {
	permissionsRole := []model.PermissionsRoles{}
	err := repo.db.Table(model.TablePermissionRole).Where("role_id = ?", roleId).Find(&permissionsRole).Error
	return permissionsRole, err
}

func (repo *repository) CheckPermission(permissionId int, namePermission string) (bool, error) {
	permissionsRole := model.Permissions{}
	err := repo.db.Table(model.TablePermission).Where("id = ?", permissionId).First(&permissionsRole).Error
	return permissionsRole.Permission == namePermission, err
}

func (repo *repository) GetAllUsers() ([]model.Users, error) {
	allUsers := []model.Users{}
	err := repo.db.Table(model.TableUserName).Find(&allUsers).Error
	return allUsers, err
}

func (repo *repository) GetRole(roleId int) (model.Roles, error) {
	role := model.Roles{}
	err := repo.db.Table(model.TableRolesName).Where("id = ?", roleId).First(&role).Error
	return role, err
}
