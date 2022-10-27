package model

type Certificate struct {
	ID      uint32 `gorm:"primaryKey;autoIncrement"`
	UserID  uint32 `gorm:"not null"`
	EventID uint32 `gorm:"not null"`
	FileURL string `gorm:"not null"`
	User    User   `gorm:"foreingKey:UserID;constraint:OnDelete:CASCADE;"`
	Event   Event  `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE;"`
}
