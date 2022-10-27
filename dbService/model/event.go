package model

import "time"

type Event struct {
	ID               uint32    `gorm:"primaryKey;autoIncrement"`
	Name             string    `gorm:"type:varchar(255);not null"`
	Description      string    `gorm:"type:varchar(255);not null"`
	CreatedAt        time.Time `gorm:"type:timestamp;default:NOW()"`
	UpdatedAt        time.Time `gorm:"type:timestamp;default:NOW()"`
	EventImg         string    `gorm:"type:varchar(255);"`
	Password         string    `gorm:"type:varchar(255);"`
	Requirement      string    `gorm:"type:varchar(255);not null"`
	NeedPayment      bool      `gorm:"not null"`
	NeedID           bool      `gorm:"not null"`
	PaymentID        uint32    `gorm:"default:null"`
	PublisherID      uint32    `gorm:"not null;default:null"`
	PaymentPrice     float64   `gorm:"default:null"`
	EventCategory    uint8     `gorm:"not null;default:null"`
	IsOffline        bool      `gorm:"not null"`
	LocationAddress  string    `gorm:"type:varchar(255);"`
	LocationCity     string    `gorm:"type:varchar(255);"`
	LocationProvince string    `gorm:"type:varchar(255);"`
	StartDate        time.Time `gorm:"not null"`
	FinishDate       time.Time `gorm:"not null"`
	Status           uint8     `gorm:"not null"`
	PresenceQuestion string    `gorm:"not null"`
	MediaLink        string    `gorm:"not null"`
	Creator          User      `gorm:"foreignKey:PublisherID;constraint:OnDelete:CASCADE;"`
}
