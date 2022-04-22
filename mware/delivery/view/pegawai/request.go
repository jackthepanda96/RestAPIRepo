package pegawai

type InsertPegawaiRequest struct {
	Nama     string `json:"nama" validate:"required"`
	HP       int    `json:"hp"`
	Password string `json:"password"`
}

type LoginRequest struct {
	HP       int    `json:"hp" validate:"required"`
	Password string `json:"password" validate:"required"`
}
