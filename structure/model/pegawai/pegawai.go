package pegawai

import (
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type Pegawai struct {
	gorm.Model
	Nama string
	HP   int
}

type PegawaiModel struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *PegawaiModel {
	return &PegawaiModel{
		Db: db,
	}
}

func (pr *PegawaiModel) Insert(newPegawai Pegawai) (Pegawai, error) {
	if err := pr.Db.Create(&newPegawai).Error; err != nil {
		log.Warn(err)
		return Pegawai{}, errors.New("tidak bisa insert data")
	}
	log.Info()
	return newPegawai, nil
}

func (pr *PegawaiModel) GetAll() ([]Pegawai, error) {
	arrPegawai := []Pegawai{}

	if err := pr.Db.Find(&arrPegawai).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select data")
	}

	if len(arrPegawai) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrPegawai, nil
}
