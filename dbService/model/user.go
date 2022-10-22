package model

import "time"

type User struct {
	ID         uint32    `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"type:varchar(255);not null"`
	Desc       string    `gorm:"type:varchar(255);not null"`
	ProfileImg string    `gorm:"type:varchar(255);"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:NOW()"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:NOW()"`
	Email      string    `gorm:"type:varchar(255);not null"`
	Password   string    `gorm:"type:text;not null"`
}
