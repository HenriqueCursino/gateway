package migration

import (
	"github.com/henriquecursino/gateway/database/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	return db.AutoMigrate(&model.Roles{}, &model.Users{}, &model.Permissions{}, &model.PermissionsRoles{})
}
