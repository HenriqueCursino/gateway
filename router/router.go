package router

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/common/errors"
	structure "github.com/henriquecursino/gateway/database"
	"github.com/henriquecursino/gateway/database/migration"
	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/database/seed"
	"gorm.io/gorm"
)

func Router() {
	router := gin.Default()
	db := structure.ConnectDB()

	if common.CurrentMode == common.DEVELOPMENT && isNeedSeed(db) {
		seed.Run(db)
	}

	router.Run(":8080")
}

func isNeedSeed(db *gorm.DB) bool {
	return existTableUsers(db) && tableUsersIsEmpty(db)
}

func existTableUsers(db *gorm.DB) bool {
	err := migration.Run(db)
	hasTable := db.Migrator().HasTable(&model.Users{})

	return errors.IsEmptyError(err) && hasTable
}

func tableUsersIsEmpty(db *gorm.DB) bool {
	err := db.First(&model.Users{}).Error

	return errors.IsEmptyError(err)
}
