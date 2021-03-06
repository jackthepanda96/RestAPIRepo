package routes

import (
	"apiex/mvc/controller/pegawai"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, pc pegawai.PegawaiController) {
	e.GET("/pegawai", pc.GetAllPegawai)
	e.POST("/pegawai", pc.InsertNewPegawai)
}
