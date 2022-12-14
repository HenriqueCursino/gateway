package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableUserName = "users"
)

type Users struct {
	ID        int    `gorm:"primaryKey; autoIncrement" `
	FullName  string `gorm:"size:30"`
	UserId    string `gorm:"unique"`
	Email     string `gorm:"size:30; unique" `
	Document  string `gorm:"unique" `
	Password  string `gorm:"size:256"`
	RoleId    int
	Roles     Roles `gorm:"foreignKey:RoleId"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
