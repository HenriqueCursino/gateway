package model

type Users struct {
	ID       int    `gorm:"primaryKey; autoIncrement" `
	FullName string `gorm:"size:30"`
	Email    string `gorm:"size:30; unique" `
	Cpf      int64  `gorm:"unique" `
	Password string `gorm:"size:15"`
	RoleId   int
	Roles    Roles `gorm:"foreignKey:RoleId"`
}
