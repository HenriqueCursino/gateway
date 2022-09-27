package router

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common"
	structure "github.com/henriquecursino/gateway/database"
	"github.com/henriquecursino/gateway/database/migration"
	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/database/seed"
	"gorm.io/gorm"
)

func Router() {
	router := gin.Default()
	db := structure.ConnectDB()

	if common.CurrentMode == common.DEVELOPMENT {
		checkNeedSeed(db)
	}

	router.Run(":8080")
}

func checkNeedSeed(db *gorm.DB) {
	if existTableUsers(db) && tableUsersIsEmpty(db) {
		seed.Run(db)
	}
}

func existTableUsers(db *gorm.DB) bool {
	err := migration.Run(db)
	hasTable := db.Migrator().HasTable(&model.Users{})

	return err == nil && hasTable
}

func tableUsersIsEmpty(db *gorm.DB) bool {
	err := db.First(&model.Users{}).Error

	return err == nil
}
