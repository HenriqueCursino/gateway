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

	router.POST("/users", middleware.Validate(), middle.CheckPermission(common.PermissionUserCreate), controller.PostUser)
	router.GET("/users", middleware.Validate(), middle.CheckPermission(common.PermissionGetUsers), controller.GetAllUsers)
	router.PUT("/users/:doc", middleware.Validate(), middle.CheckPermission(common.PermissionUserUpdate), controller.UpdateUserRole)
	router.DELETE("/users", middleware.Validate(), middle.CheckPermission(common.PermissionUserDelete), controller.DeleteUser)

	router.POST("/roles", middleware.Validate(), middle.CheckPermission(common.PermissionRoleCreate), controller.PostRole)
	router.GET("/roles", middleware.Validate(), middle.CheckPermission(common.PermissionGetRole), controller.GetAllRoles)
	router.DELETE("/roles", middleware.Validate(), middle.CheckPermission(common.PermissionRoleDelete), controller.DeleteRole)

	router.POST("/login", controller.Login)

	router.Run(os.Getenv("SERVER_PORT"))
}
