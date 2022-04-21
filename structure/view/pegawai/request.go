package pegawai

type InsertPegawaiRequest struct {
	Nama string `json:"nama" validate:"required"`
	HP   int    `json:"hp"`
	Gaji int32  `json:"gaji"`
}
