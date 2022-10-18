package user

import (
	"time"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Dob       time.Time `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Username  string    `gorm:"unique;not null"`
	Sosmed    []sosmed.Sosmed
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
