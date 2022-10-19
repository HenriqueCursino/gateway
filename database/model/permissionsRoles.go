package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	TablePermissionRole = "permissions_roles"
)

type PermissionsRoles struct {
	ID           int `gorm:"primaryKey; autoIncrement"`
	RoleId       int
	Roles        Roles `gorm:"foreignKey:RoleId"`
	PermissionId int
	Permissions  Permissions `gorm:"foreignKey:PermissionId"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
