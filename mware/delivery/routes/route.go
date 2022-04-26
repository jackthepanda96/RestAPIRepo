package routes

import (
	"apiex/mware/delivery/controller/book"
	"apiex/mware/delivery/controller/pegawai"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, pc pegawai.ControllerPegawai, bc book.ControllerBook) {
	// e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	// e.GET("/pegawai", pc.GetAllPegawai)
	e.POST("/pegawai", pc.Insert) // Register
	e.POST("/login", pc.Login)    // Login
	e.GET("/coba", pc.GetAllPegawai(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RH$SI4")}))
	e.POST("/book", bc.Insert, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RH$SI4")}))
	// kelompokGET := e.Group("/pegawai")
	// kelompokGET.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	// }))
	// kelompokGET.GET("/coba", pc.GetAllPegawai, middleware.BasicAuth(mwareFunction.BasicCheck))

}
