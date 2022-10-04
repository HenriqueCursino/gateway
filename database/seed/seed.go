package seed

import (
	"fmt"

	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/tools"
	"gorm.io/gorm"
)

type types interface {
	*[]model.Roles | *[]model.Users | *[]model.Permissions | *[]model.PermissionsRoles
}

func Run(db *gorm.DB) {
	roles := []model.Roles{
		{
			ID:   1,
			Role: "admin",
		},
		{
			ID:   2,
			Role: "commun",
		},
	}

	permissions := []model.Permissions{
		{
			ID:         1,
			Permission: "user_create",
		},
		{
			ID:         2,
			Permission: "user_delete",
		},
		{
			ID:         3,
			Permission: "user_update",
		},
		{
			ID:         4,
			Permission: "user_read",
		},
		{
			ID:         5,
			Permission: "permission_create",
		},
		{
			ID:         6,
			Permission: "permission_delete",
		},
		{
			ID:         7,
			Permission: "permission_update",
		},
		{
			ID:         8,
			Permission: "permission_read",
		},
	}

	users := []model.Users{
		{
			FullName: "Henrique Cursino",
			Hash:     tools.GenerateHash(),
			Email:    "henrique@gmail.com",
			Document: "12345678910",
			Password: "123",
			RoleId:   1,
			Roles:    model.Roles{},
		},
		{
			FullName: "Guilherme Sembeneli",
			Hash:     tools.GenerateHash(),
			Email:    "guilherme@gmail.com",
			Document: "11122233344",
			Password: "123456",
			RoleId:   2,
			Roles:    model.Roles{},
		},
	}

	permissionsRoles := []model.PermissionsRoles{
		{
			RoleId:       roles[0].ID,
			PermissionId: permissions[0].ID,
		},
		{
			RoleId:       roles[0].ID,
			PermissionId: permissions[1].ID,
		},
		{
			RoleId:       roles[0].ID,
			PermissionId: permissions[2].ID,
		},
		{
			RoleId:       roles[0].ID,
			PermissionId: permissions[3].ID,
		},
		{
			RoleId:       roles[0].ID,
			PermissionId: permissions[4].ID,
		},
		{
			RoleId:       roles[0].ID,
			PermissionId: permissions[5].ID,
		},
		{
			RoleId:       roles[0].ID,
			PermissionId: permissions[6].ID,
		},
		{
			RoleId:       roles[0].ID,
			PermissionId: permissions[7].ID,
		},
	}

	create(db, &roles)
	create(db, &permissions)
	create(db, &users)
	create(db, &permissionsRoles)
}

func create[model types](db *gorm.DB, seeds model) {
	if err := db.Create(&seeds).Error; err != nil {
		fmt.Println(err.Error())
	}
}
