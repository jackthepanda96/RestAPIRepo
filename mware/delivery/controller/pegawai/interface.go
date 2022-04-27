package pegawai

import "github.com/labstack/echo/v4"

type ControllerPegawai interface {
	Insert() echo.HandlerFunc
	GetAllPegawai() echo.HandlerFunc
	Login() echo.HandlerFunc
}
