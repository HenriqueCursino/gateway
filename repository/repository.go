package repository

import (
	"errors"

	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/dto"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *dto.UserCreate) error
	LoginUser(login dto.UserLogin) (model.Users, error)
	GetUser(hash string) (model.Users, error)
	GetAllPermissionsRole(roleId int) ([]model.PermissionsRoles, error)
	CheckPermissionRepository(permissionId int, namePermission string) (bool, error)
	GetAllUsers() ([]model.Users, error)
	GetRole(roleId int) (model.Roles, error)
	UpdateUserRole(document string, newRoleId int) error
	DeleteUser(userId string) error
	CreateRole(role dto.RoleUser) error
	GetAllRoles() []model.Roles
	GetAllPermissions(roleId int) model.Permissions
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
	var user model.Users
	err := repo.db.Table(model.TableUserName).Where("email = ?", login.Email).First(&user).Error
	return user, err
}

func (repo *repository) GetUser(hash string) (model.Users, error) {
	var user model.Users
	err := repo.db.Table(model.TableUserName).Where("user_id = ?", hash).First(&user).Error
	return user, err
}

func (repo *repository) GetAllPermissionsRole(roleId int) ([]model.PermissionsRoles, error) {
	var permissionsRole []model.PermissionsRoles
	err := repo.db.Table(model.TablePermissionRole).Where("role_id = ?", roleId).Find(&permissionsRole).Error
	return permissionsRole, err
}

func (repo *repository) CheckPermissionRepository(permissionId int, namePermission string) (bool, error) {
	var permissionsRole model.Permissions
	err := repo.db.Table(model.TablePermission).Where("id = ?", permissionId).First(&permissionsRole).Error
	return permissionsRole.Permission == namePermission, err
}

func (repo *repository) GetAllUsers() ([]model.Users, error) {
	var allUsers []model.Users
	err := repo.db.Table(model.TableUserName).Find(&allUsers).Error
	return allUsers, err
}

func (repo *repository) GetRole(roleId int) (model.Roles, error) {
	var role model.Roles
	err := repo.db.Table(model.TableRolesName).Where("id = ?", roleId).First(&role).Error
	return role, err
}

func (repo *repository) UpdateUserRole(document string, newRoleId int) error {
	err := repo.db.Table(model.TableUserName).
		Where("document = ?", document).
		Update("role_id", newRoleId).
		Error

	return err
}

func (repo *repository) DeleteUser(userId string) error {
	var userDeleted model.Users
	query := repo.db.Table(model.TableUserName).Where("user_id = ?", userId).Delete(&userDeleted)
	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected < 1 {
		return errors.New("dont have user in database")
	}

	return nil
}

func (repo *repository) CreateRole(role dto.RoleUser) error {
	roleModel := model.Roles{
		Role: role.Role,
	}
	result := repo.db.Table(model.TableRolesName).Create(&roleModel)
	if result.Error != nil {
		return result.Error
	}

	newRole := result.Statement.Model.(*model.Roles)
	for i := 0; i < len(role.PermissionsId); i++ {
		addPermission := dto.PermissionRole{
			RoleId:       newRole.ID,
			PermissionId: role.PermissionsId[i],
		}
		repo.db.Table(model.TablePermissionRole).Create(&addPermission)
	}
	return nil
}

func (repo *repository) GetAllRoles() []model.Roles {
	var allRoles []model.Roles
	repo.db.Table(model.TableRolesName).Find(&allRoles)
	return allRoles
}

func (repo *repository) GetAllPermissions(permissionId int) model.Permissions {
	var permissions model.Permissions
	repo.db.Table(model.TablePermission).Where("id = ?", permissionId).First(&permissions)
	return permissions
}
