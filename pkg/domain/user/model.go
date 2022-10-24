package user

import (
	"time"

	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/comment"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/photo"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"
	"gorm.io/gorm"
)

type User struct {
	ID        uint            `gorm:"primaryKey"`
	Dob       time.Time       `gorm:"not null"`
	Email     string          `gorm:"unique;not null"`
	Password  string          `gorm:"not null"`
	Username  string          `gorm:"unique;not null"`
	Sosmed    []sosmed.Sosmed `gorm:"constraint:OnDelete:CASCADE"`
	Photos    []photo.Photo   `gorm:"constraint:OnDelete:CASCADE"`
	Comments  []comment.Comment
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
