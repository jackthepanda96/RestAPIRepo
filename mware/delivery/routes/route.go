package routes

import (
	"apiex/mware/delivery/controller"
	mwareFunction "apiex/mware/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, pc controller.ControllerPegawai) {
	// e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.RemoveTrailingSlash())

	// e.GET("/pegawai", pc.GetAllPegawai)
	e.POST("/pegawai", pc.Insert)

	kelompokGET := e.Group("/pegawai")
	kelompokGET.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	kelompokGET.GET("/coba", pc.GetAllPegawai, middleware.BasicAuth(mwareFunction.BasicCheck))

}
