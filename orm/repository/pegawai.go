package repository

import (
	"apiex/orm/entities"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PegawaiRepo struct {
	Db *gorm.DB
}

func (pr *PegawaiRepo) Insert(newPegawai entities.Pegawai) (entities.Pegawai, error) {
	if err := pr.Db.Create(&newPegawai).Error; err != nil {
		log.Warn(err)
		return entities.Pegawai{}, errors.New("tidak bisa insert data")
	}
	log.Info()
	return newPegawai, nil
}

func (pr *PegawaiRepo) GetAll() ([]entities.Pegawai, error) {
	arrPegawai := []entities.Pegawai{}

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
