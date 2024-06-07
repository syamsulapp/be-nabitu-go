package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	Fullname   string         `json:"fullname" form:"fullname" validate:"required,min=6" gorm:"not null"`
	Age        int            `json:"age" form:"age" validate:"required,number" gorm:"not null"`
	Alamat     string         `json:"alamat" form:"alamat" validate:"required,min=6" gorm:"not null"`
	Users_id   int            `json:"users_id" form:"users_id" validate:"required,number" gorm:"not null"`
	Created_at string         `json:"created_at"`
	Updated_at string         `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
