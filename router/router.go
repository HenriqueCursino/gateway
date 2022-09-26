package router

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/structure"
)

func Router() {
	router := gin.Default()
	db := structure.ConnectDB()

	db.AutoMigrate(&model.Roles{})
	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Permissions{})
	db.AutoMigrate(&model.PermissionsRoles{})

	router.Run(":8080")
}
