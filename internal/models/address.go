package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID uint64 `gorm:"primaryKey;autoIncrement"`

	UserID uint64 `gorm:"index;not null"`

	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(255);index;not null"`
	Phone     string `gorm:"type:varchar(20)" validate:"omitempty"`

	AddressLine1 string `gorm:"type:varchar(255);not null"`
	AddressLine2 string `gorm:"type:varchar(255)"`
	City         string `gorm:"type:varchar(100)"`
	State        string `gorm:"type:varchar(100)"`
	Country      string `gorm:"type:varchar(100)"`
	Pincode      string `gorm:"type:varchar(20)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CreateAddressRequest struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" validate:"required, strict_email"` 
	Phone        string `json:"phone" validate:"phone"`                     
	AddressLine1 string `json:"address_line1" validate:"required"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Pincode      string `json:"pincode" validate:"pincode"` 
}

type UpdateAddressRequest struct {
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	Email        *string `json:"email" validate:"strict_email"` 
	Phone        *string `json:"phone" validate:"phone"` 
	AddressLine1 *string `json:"address_line1" validate:"required"`
	AddressLine2 *string `json:"address_line2"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	Country      *string `json:"country"`
	Pincode      *string `json:"pincode" validate:"pincode"` 
}

type ExportAddressRequest struct {
	Fields []string `json:"fields" validate:"required,min=1"`
	Email  string   `json:"email" validate:"required, strict_email"` 
}

type AddressResponse struct {
	Id           uint64 `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Pincode      string `json:"pincode"`
}

type ListAddressQuery struct {
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
	Search  string `form:"search"`
	City    string `form:"city"`
	State   string `form:"state"`
	Country string `form:"country"`
	Pincode string `form:"pincode"`
}
