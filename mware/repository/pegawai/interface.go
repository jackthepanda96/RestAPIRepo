package pegawai

import "apiex/mware/entity"

type Pegawai interface {
	Insert(newPegawai entity.Pegawai) (entity.Pegawai, error)
	GetAll() ([]entity.Pegawai, error)
}
