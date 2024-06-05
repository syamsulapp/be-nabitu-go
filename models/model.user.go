package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	Password   []byte         `json:"password"`
	Created_at string         `json:"created_at"`
	Updated_at string         `json:"updated_at`
	Deleted_at gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
