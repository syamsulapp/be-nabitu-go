package schemas

import (
	"time"

	"gorm.io/gorm"
)

type SchemaIndexProfile struct {
	Message string `json:"mesage"`
}

type Profile struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Fullname  string         `json:"name" form:"name"`
	Age       string         `json:"age" form:"age"`
	Alamat    string         `json:"alamat" form:"Alamat"`
	UsersId   int            `json:"users_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
