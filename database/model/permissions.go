package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	TablePermission = "permissions"
)

type Permissions struct {
	ID         int    `gorm:"primaryKey; autoIncrement"`
	Permission string `gorm:"size:30"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
