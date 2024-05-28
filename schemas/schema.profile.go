package schemas

type SchemaIndexProfile struct {
	Message string `json:"mesage"`
}

type SchemaDatabaseError struct {
	Type string
	Code int
}

type Profiles struct {
	ID       string `json:"id"`
	Fullname string `json:"name"`
	Age      string `json:"age"`
	Alamat   string `json:"alamat"`
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
