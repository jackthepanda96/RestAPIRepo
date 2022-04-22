package pegawai

import "apiex/layered/entity"

type Pegawai interface {
	Insert(newPegawai entity.Pegawai) (entity.Pegawai, error)
	GetAll() ([]entity.Pegawai, error)
}
