package model

import "time"

type Register struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	UserID      uint32 `gorm:"not null"`
	EventID     uint32 `gorm:"not null"`
	Requirement string `gorm:"type:text;not null"`
	PaymentImg  string
	IDImg       string
	CreatedAt   time.Time `gorm:"type:timestamp;default:NOW()"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Event       Event     `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE;"`
}
