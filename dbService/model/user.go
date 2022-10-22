package model

import "time"

type User struct {
	ID         uint32    `gorm:""`
	Name       string    `gorm:"not null"`
	Desc       string    `gorm:"not null"`
	ProfileImg string    `gorm:""`
	CreatedAt  time.Time `gorm:""`
	UpdatedAt  time.Time `gorm:""`
	Email      string    `gorm:"not null"`
	Password   string    `gorm:"not null"`
}
