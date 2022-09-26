package migration

import (
	"github.com/henriquecursino/gateway/database/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Roles{})
	err = db.AutoMigrate(&model.Users{})
	err = db.AutoMigrate(&model.Permissions{})
	err = db.AutoMigrate(&model.PermissionsRoles{})
	return err
}
