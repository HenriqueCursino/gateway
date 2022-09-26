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
			Role: "ADM",
		},
		{
			Role: "Commun",
		},
	}

	permissions := []model.Permissions{
		{
			Permission: "create_user",
		},
		{
			Permission: "delete_user",
		},
		{
			Permission: "update_user_permission",
		},
		{
			Permission: "check_user_permission",
		},
		{
			Permission: "delete_user",
		},
		{
			Permission: "delete_user_permission",
		},
	}

	create(db, &roles)
	create(db, &permissions)
}

func create[model types](db *gorm.DB, seeds model) {
	if err := db.Create(&seeds).Error; err != nil {
		fmt.Println(err.Error())
	}
}
