package model

type PermissionsRoles struct {
	ID           int `gorm:"primaryKey; autoIncrement"`
	RoleId       int
	Roles        Roles `gorm:"foreignKey:RoleId"`
	PermissionId int
	Permissions  Permissions `gorm:"foreignKey:PermissionId"`
}
