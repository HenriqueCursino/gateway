package model

const (
	TableUserName = "users"
)

type Users struct {
	ID       int    `gorm:"primaryKey; autoIncrement" `
	FullName string `gorm:"size:30"`
	Hash     string `gorm:"unique"`
	Email    string `gorm:"size:30; unique" `
	Document string `gorm:"unique" `
	Password string `gorm:"size:15"`
	RoleId   int
	Roles    Roles `gorm:"foreignKey:RoleId"`
}
