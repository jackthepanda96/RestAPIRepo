package pegawai

import (
	"apiex/layered/entity"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PegawaiRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *PegawaiRepo {
	return &PegawaiRepo{
		Db: db,
	}
}

func (pr *PegawaiRepo) Insert(newPegawai entity.Pegawai) (entity.Pegawai, error) {
	if err := pr.Db.Create(&newPegawai).Error; err != nil {
		log.Warn(err)
		return entity.Pegawai{}, errors.New("tidak bisa insert data")
	}
	log.Info()
	return newPegawai, nil
}

func (pr *PegawaiRepo) GetAll() ([]entity.Pegawai, error) {
	arrPegawai := []entity.Pegawai{}

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
