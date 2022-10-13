package tools

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/common/errors"
	"github.com/henriquecursino/gateway/database/migration"
	"github.com/henriquecursino/gateway/database/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RemoveMask(document string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return re.ReplaceAllString(document, "")
}

func GenerateHash() string {
	return uuid.New().String()
}

func GetStringFromBody(body interface{}) string {
	return fmt.Sprintf("%v", body)
}

func IsNeedSeed(db *gorm.DB) bool {
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

func Encrypt(hash string) (*string, error) {
	byteHash := []byte(hash)

	hashedPassword, err := bcrypt.GenerateFromPassword(byteHash, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	encryptedHash := string(hashedPassword)
	return &encryptedHash, nil
}

func Compare(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
