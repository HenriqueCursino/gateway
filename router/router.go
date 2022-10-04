package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/common/env"
	"github.com/henriquecursino/gateway/common/errors"
	"github.com/henriquecursino/gateway/controller"
	dataBase "github.com/henriquecursino/gateway/database"
	"github.com/henriquecursino/gateway/database/migration"
	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/database/seed"
	"github.com/henriquecursino/gateway/middleware"
	"github.com/henriquecursino/gateway/repository"
	"github.com/henriquecursino/gateway/service"
	"gorm.io/gorm"
)

func Router() {
	router := gin.Default()
	db := dataBase.ConnectDB()
	if env.CurrentMode == common.Development && isNeedSeed(db) {
		seed.Run(db)
	}

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	controller := controller.NewController(serv)

	router.POST("/users", middleware.Validate(), controller.PostUser)
	router.POST("/login", controller.Login)

	router.Run(os.Getenv("SERVER_PORT"))
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
	var userModel []model.Users
	err := db.Find(&userModel)

	return err != nil && len(userModel) < common.CheckTableEmpty
}
