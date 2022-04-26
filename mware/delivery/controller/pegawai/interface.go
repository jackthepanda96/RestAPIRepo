package pegawai

import "github.com/labstack/echo/v4"

type ControllerPegawai interface {
	Insert(c echo.Context) error
	GetAllPegawai() echo.HandlerFunc
	Login(c echo.Context) error
}
