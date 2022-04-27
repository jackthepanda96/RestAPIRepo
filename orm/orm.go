package main

import (
	"apiex/orm/controller"
	"apiex/orm/entities"
	"apiex/orm/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:password*@tcp(host:3306)/be8rds?charset=utf8mb4&parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

var arrPegawai []entities.Pegawai

func InsertPegawai(c echo.Context) error {
	newPegawai := entities.Pegawai{}

	if err := c.Bind(&newPegawai); err != nil {
		log.Warn(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "cannot bind data",
			"data":    nil,
		})
	}

	arrPegawai = append(arrPegawai, newPegawai)

	log.Info()
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "success create data",
		"data":    newPegawai,
	})
}

func GetPegawai(c echo.Context) error {
	param1 := c.QueryParam("category")
	fmt.Println(param1)

	if len(arrPegawai) == 0 {
		log.Warn("no data found")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "connot find data",
			"data":    nil,
		})
	}
	log.Info()
	return c.JSON(200, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success find data",
		"data":    arrPegawai,
	})
}

func GetPegawaiByID(c echo.Context) error {
	id := c.Param("id")

	convID, err := strconv.Atoi(id)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "connot convert ID",
			"data":    nil,
		})
	}

	if convID >= len(arrPegawai) {
		log.Error("Index out of Range")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Index out of Range",
			"data":    nil,
		})
	}
	log.Info()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Found Data",
		"data":    arrPegawai[convID],
	})

}

func main() {

	e := echo.New()
	db := InitDB()
	db.AutoMigrate(&entities.Pegawai{})
	repo := repository.PegawaiRepo{db}
	controll := controller.PegawaiController{repo}

	e.GET("/halo", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Halo ini via echo")
	})
	e.GET("/pegawai", controll.GetAllPegawai)
	e.POST("/pegawai", controll.InsertNewPegawai)
	// e.GET("/pegawai/:id", GetPegawaiByID) // pegawai/1 | pegawai/2

	e.Start(":8000")

}
