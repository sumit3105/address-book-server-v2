package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint64 `gorm:"primaryKey;autoIncrement"`

	Email    string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Addresses []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,strict_email"` 
	Password string `json:"password" validate:"required"` 
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,strict_email"` 
	Password string `json:"password" validate:"required"` 
}
