package main

import (
	"apiex/mware/config"
	pegawaiController "apiex/mware/delivery/controller"
	"apiex/mware/delivery/controller/book"
	"apiex/mware/delivery/routes"
	"apiex/mware/entity"
	bookRepo "apiex/mware/repository/book"
	pegawaiRepo "apiex/mware/repository/pegawai"
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	db.AutoMigrate(entity.Pegawai{})
	db.AutoMigrate(entity.Book{})
	e := echo.New()

	repoUser := pegawaiRepo.New(db)
	repoBook := bookRepo.New(db)

	controllUser := pegawaiController.New(repoUser, validator.New())
	controllBook := book.New(repoBook, validator.New())

	routes.RegisterPath(e, controllUser, controllBook)

	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
