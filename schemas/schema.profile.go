package schemas

import "gorm.io/gorm"

type SchemaIndexProfile struct {
	Message string `json:"mesage"`
}

type Profile struct {
	gorm.Model
	nama_depan        string
	nama_belakang     string
	nama_panggilan    string
	jenis_kelamin     string
	tanggal_lahir     string
	tempat_lahir      string
	umur              string
	status_pernikahan string
	domisili          string
	alamat_rumah      string
	nomor_handphone   string
	nomor_telpon      string
}
