package schemas

import (
	"time"

	"gorm.io/gorm"
)

type SchemaIndexProfile struct {
	Message string `json:"mesage"`
}

type SchemaDatabaseError struct {
	Type string
	Code int
}

type Profile struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Fullname  string         `json:"name" form:"name"`
	Age       string         `json:"age" form:"age"`
	Alamat    string         `json:"alamat" form:"Alamat"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Jenis_Kelamin     string `json:"jenis_kelamin"`
	// Tanggal_Lahir     string `json:"tanggal_lahir"`
	// Tempat_Lahir      string `json:"tempat_lahir"`
	// Umur              string `json:"umur"`
	// Status_Pernikahan string `json:"status_pernikahan"`
	// Domisili          string `json:"domisili"`
	// Alamat_Rumah      string `json:"alamat_rumah"`
	// Nomor_Handphone   string `json:"nomor_handphone"`
	// Nomor_Telpon      string `json:"nomor_telpon"`
}
