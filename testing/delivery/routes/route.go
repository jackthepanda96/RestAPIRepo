package routes

import (
	"apiex/testing/delivery/controller"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, pc controller.ControllerPegawai) {
	// e.GET("/pegawai", pc.GetAllPegawai)
	e.POST("/pegawai", pc.Insert)
}
