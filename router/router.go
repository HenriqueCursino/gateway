package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/common/env"
	"github.com/henriquecursino/gateway/controller"
	dataBase "github.com/henriquecursino/gateway/database"
	"github.com/henriquecursino/gateway/database/seed"
	"github.com/henriquecursino/gateway/middleware"
	"github.com/henriquecursino/gateway/repository"
	"github.com/henriquecursino/gateway/service"
	"github.com/henriquecursino/gateway/tools"
)

func Router() {
	router := gin.Default()
	db := dataBase.ConnectDB()
	if env.CurrentMode == common.Development && tools.IsNeedSeed(db) {
		seed.Run(db)
	}

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	middle := middleware.NewMiddleware(repo)
	controller := controller.NewController(serv, middle)

	router.POST("/users", middleware.Validate(), controller.PostUser)
	router.GET("/users", middleware.Validate(), controller.GetAllUsers)
	router.PUT("/users/:doc", middleware.Validate(), controller.UpdateUserRole)
	router.DELETE("/user", middleware.Validate(), controller.DeleteUser)
	router.POST("/login", controller.Login)

	router.Run(os.Getenv("SERVER_PORT"))
}
