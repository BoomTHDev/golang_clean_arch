package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	ADMIN UserRole = "admin"
	USER  UserRole = "user"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primarykey;"`
	Name      string    `gorm:"not null;"`
	Email     string    `gorm:"not null;unique;"`
	Password  string    `gorm:"not null;"`
	Role      UserRole  `gorm:"not null;default:user"`
	CreatedAt time.Time `gorm:"autoCreateTime;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"`
}
