package model

type PermissionsRoles struct {
	ID           int         `gorm:"primaryKey; autoIncrement" json:"id"`
	RoleId       int         `json:"id_role"`
	Roles        Roles       `gorm:"foreignKey:RoleId"`
	PermissionId int         `json:"id_permission"`
	Permissions  Permissions `gorm:"foreignKey:PermissionId"`
}
