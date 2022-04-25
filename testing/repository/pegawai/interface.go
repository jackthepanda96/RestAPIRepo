package pegawai

import "apiex/testing/entity"

type Pegawai interface {
	Insert(newPegawai entity.Pegawai) (entity.Pegawai, error)
	GetAll() ([]entity.Pegawai, error)
}
