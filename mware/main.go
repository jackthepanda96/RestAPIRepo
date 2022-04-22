package main

import (
	"apiex/layered/config"
	pegawaiController "apiex/layered/delivery/controller"
	"apiex/layered/delivery/routes"
	pegawaiRepo "apiex/layered/repository/pegawai"
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	e := echo.New()

	repo := pegawaiRepo.New(db)
	controller := pegawaiController.New(repo, validator.New())

	routes.RegisterPath(e, controller)

	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
