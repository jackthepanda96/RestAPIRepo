package controller

import "github.com/labstack/echo/v4"

type ControllerPegawai interface {
	Insert(c echo.Context) error
	GetAllPegawai(c echo.Context) error
	Login(c echo.Context) error
}
