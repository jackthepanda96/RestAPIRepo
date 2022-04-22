package entity

import "gorm.io/gorm"

type Pegawai struct {
	gorm.Model
	Nama string
	HP   int
}
