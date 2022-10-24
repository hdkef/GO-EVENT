package model

type Event struct {
	ID               uint32 `gorm:"primaryKey;autoIncrement"`
	Name             string `gorm:"type:varchar(255);not null"`
	Desc             string `gorm:"type:varchar(255);not null"`
	EventImg         string `gorm:"type:varchar(255);"`
	Password         string `gorm:"type:varchar(255);"`
	Requirement      string `gorm:"type:varchar(255);not null"`
	NeedPayment      bool   `gorm:"not null"`
	NeedID           bool   `gorm:"not null"`
	PaymentID        uint32
	CreatorID        uint32 `gorm:"not null"`
	PaymentPrice     float64
	EventCategory    uint8  `gorm:"not null"`
	IsOffline        bool   `gorm:"not null"`
	LocationAddress  string `gorm:"type:varchar(255);"`
	LocationCity     string `gorm:"type:varchar(255);"`
	LocationProvince string `gorm:"type:varchar(255);"`
	StartDate        string `gorm:"not null"`
	FinishDate       string `gorm:"not null"`
	Status           uint8  `gorm:"type:varchar(255);"`
	PresenceQuestion string `gorm:"not null"`
	MediaLink        string `gorm:"not null"`
	Creator          User   `gorm:"foreignKey:creatorID;constraint:OnDelete:CASCADE;"`
}
