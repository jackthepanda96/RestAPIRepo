package main

import (
	"apiex/mvc/config"
	"apiex/mvc/controller/pegawai"
	mPegawai "apiex/mvc/model/pegawai"
	"apiex/mvc/routes"
	"fmt"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	e := echo.New()
	// e.Validator = validator.New()
	// pm := mPegawai.PegawaiModel{Db: db}
	pm := mPegawai.New(db)
	pc := pegawai.PegawaiController{Repo: *pm, Valid: validator.New()}

	routes.RegisterPath(e, pc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))

}
