package pegawai

import (
	"apiex/layered/entity"
	"net/http"
)

func SuccessInsert(data entity.Pegawai) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data pegawai",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data pegawai",
		"status":  false,
		"data":    nil,
	}
}
