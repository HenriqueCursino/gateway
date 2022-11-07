package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableRolesName = "roles"
)

type Roles struct {
	ID        int `gorm:"primaryKey; autoIncrement"`
	Role      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
