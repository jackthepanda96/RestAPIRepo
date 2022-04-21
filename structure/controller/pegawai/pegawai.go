package pegawai

import (
	mPegawai "apiex/structure/model/pegawai"
	"apiex/structure/view"
	"apiex/structure/view/pegawai"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PegawaiController struct {
	Repo mPegawai.PegawaiModel
}

func (pc *PegawaiController) InsertNewPegawai(c echo.Context) error {
	var tmpPegawai pegawai.InsertPegawaiRequest

	if err := c.Bind(&tmpPegawai); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, pegawai.BadRequest())
	}
	newPegawai := mPegawai.Pegawai{Nama: tmpPegawai.Nama}
	res, err := pc.Repo.Insert(newPegawai)

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil insert")
	return c.JSON(http.StatusCreated, pegawai.SuccessInsert(res))
}

func (pc *PegawaiController) GetAllPegawai(c echo.Context) error {

	res, err := pc.Repo.GetAll()

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil get all data")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil get all data",
		"status":  true,
		"data":    res,
	})
}
