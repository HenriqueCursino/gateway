package model

type Permissions struct {
	ID         int    `gorm:"primaryKey; autoIncrement" json:"id"`
	Permission string `gorm:"size:30" json:"permission"`
}
