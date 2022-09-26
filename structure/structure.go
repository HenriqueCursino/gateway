package structure

import (
	"fmt"

	"github.com/henriquecursino/gateway/database/seed"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "henrique:teste@tcp(127.0.0.1:3306)/permissions_control?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect DataBase", err)
	}

	seed.Run(db)

	return db
}
