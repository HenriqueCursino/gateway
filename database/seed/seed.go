package seed

import (
	"fmt"

	"github.com/henriquecursino/gateway/database/model"
	"gorm.io/gorm"
)

type types interface {
	*[]model.Roles | *[]model.Users | *[]model.Permissions | *[]model.PermissionsRoles
}

func Run(db *gorm.DB) {
	roles := []model.Roles{
		{
			Role: "admin",
		},
		{
			Role: "commun",
		},
	}

	permissions := []model.Permissions{
		{
			Permission: "user_create",
		},
		{
			Permission: "user_delete",
		},
		{
			Permission: "user_update",
		},
		{
			Permission: "user_read",
		},
		{
			Permission: "permission_create",
		},
		{
			Permission: "permission_delete",
		},
		{
			Permission: "permission_update",
		},
		{
			Permission: "permission_read",
		},
	}

	users := []model.Users{
		{
			FullName: "Henrique Cursino",
			Email:    "henrique@gmail.com",
			Cpf:      12345678910,
			Password: "123",
			RoleId:   1,
			Roles:    model.Roles{},
		},
		{
			FullName: "Guilherme Sembeneli",
			Email:    "guilherme@gmail.com",
			Cpf:      11122233344,
			Password: "123456",
			RoleId:   2,
			Roles:    model.Roles{},
		},
	}

	create(db, &roles)
	create(db, &permissions)
	create(db, &users)
}

func create[model types](db *gorm.DB, seeds model) {
	if err := db.Create(&seeds).Error; err != nil {
		fmt.Println(err.Error())
	}
}
