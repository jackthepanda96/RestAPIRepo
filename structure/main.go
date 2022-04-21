package main

import (
	"apiex/structure/config"
	"apiex/structure/controller/pegawai"
	mPegawai "apiex/structure/model/pegawai"
	"apiex/structure/routes"
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
	pm := mPegawai.PegawaiModel{Db: db}
	pc := pegawai.PegawaiController{Repo: pm, Valid: validator.New()}

	routes.RegisterPath(e, pc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))

}
