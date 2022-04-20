package entities

import "gorm.io/gorm"

type Pegawai struct {
	gorm.Model
	Nama string `json:"nama"`
	HP   int    `json:"hp"`
}
