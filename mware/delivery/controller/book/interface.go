package book

import "github.com/labstack/echo/v4"

type ControllerBook interface {
	Insert(c echo.Context) error
}
