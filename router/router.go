package router

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/database/migration"
	"github.com/henriquecursino/gateway/database/model"
	"github.com/henriquecursino/gateway/database/seed"
	"github.com/henriquecursino/gateway/structure"
	"gorm.io/gorm"
)

func Router() {
	router := gin.Default()
	db := structure.ConnectDB()

	if os.Getenv("CURRENT_MODE") == common.DEVELOPMENT {
		checkNeedSeed(db)
	}

	router.Run(":8080")
}

func checkNeedSeed(db *gorm.DB) {
	if err := migration.Run(db); err == nil && db.Migrator().HasTable(&model.Users{}) {
		if err := db.First(&model.Users{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seed.Run(db)
		}
	}
}
