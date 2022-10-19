package comment

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	PhotoID   uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
