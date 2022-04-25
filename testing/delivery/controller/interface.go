package controller

import "github.com/labstack/echo/v4"

type ControllerPegawai interface {
	Insert(c echo.Context) error
}
