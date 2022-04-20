package controller

import (
	"apiex/orm/entities"
	"apiex/orm/repository"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PegawaiController struct {
	Repo repository.PegawaiRepo
}

func (pc *PegawaiController) InsertNewPegawai(c echo.Context) error {
	var tmpPegawai entities.Pegawai

	if err := c.Bind(&tmpPegawai); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "terdapat kesalahan pada input",
			"status":  false,
			"data":    nil,
		})
	}

	res, err := pc.Repo.Insert(tmpPegawai)

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "terdapat kesalahan pada server",
			"status":  false,
			"data":    nil,
		})
	}
	log.Info("berhasil insert")
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data",
		"status":  true,
		"data":    res,
	})
}

func (pc *PegawaiController) GetAllPegawai(c echo.Context) error {

	res, err := pc.Repo.GetAll()

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "terdapat kesalahan pada server",
			"status":  false,
			"data":    nil,
		})
	}
	log.Info("berhasil get all data")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil get all data",
		"status":  true,
		"data":    res,
	})
}
