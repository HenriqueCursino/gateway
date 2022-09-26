package structure

import (
	"fmt"

	"github.com/henriquecursino/gateway/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "henrique:teste@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect DataBase", err)
	}

	fmt.Println("teste")

	db.AutoMigrate(&model.Roles{})
	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Permissions{})
	db.AutoMigrate(&model.PermissionsRoles{})

	fmt.Println("Connect sucess!")

	return db
}
