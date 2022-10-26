package model

type Subscription struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	ConsumerID  uint32 `gorm:"not null"`
	PublisherID uint32 `gorm:"not null"`
	Consumer    User   `gorm:"foreignKey:ConsumerID;constraint:OnDelete:CASCADE;"`
	Publisher   User   `gorm:"foreignKey:PublisherID;constraint:OnDelete:CASCADE;"`
}
