package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	Password   []byte         `json:"password"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at`
	Deleted_at gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
