package main

import (
	"apiex/structure/config"
	"apiex/structure/controller/pegawai"
	mPegawai "apiex/structure/model/pegawai"
	"apiex/structure/routes"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	e := echo.New()
	pm := mPegawai.PegawaiModel{Db: db}
	pc := pegawai.PegawaiController{Repo: pm}

	routes.RegisterPath(e, pc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))

}
