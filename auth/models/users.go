package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserName         string         `gorm:"uniqueIndex;not null"`
	FirstName        string         `gorm:"not null"`
	Email            string         `gorm:"uniqueIndex;not null"`
	PasswordHash     *string        `gorm:""`
	CreatedAt        *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt        *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	AuthProviderType string         `gorm:"not null;default:'local'"`
	RoleID           uint           `gorm:"not null"`
	Role             Role           `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DOB              *string        `gorm:""`
	PhoneNumber      *string        `gorm:""`
	Verified         bool           `gorm:"default:false"`
	IsBlocked        bool           `gorm:"default:false"`
}

type UserAddress struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	AddressLine string `gorm:"not null"`
	City        string `gorm:"not null"`
	State       string
	Country     string `gorm:"not null"`
	ZipCode     string
}

type Role struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"`
}
