package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Pegawai struct {
	Nama string `json:"nama"`
	HP   int    `json:"hp"`
}

var arrPegawai []Pegawai

func InsertPegawai(c echo.Context) error {
	newPegawai := Pegawai{}

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

	e.GET("/halo", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Halo ini via echo")
	})
	e.GET("/pegawai", GetPegawai)
	e.POST("/pegawai", InsertPegawai)
	e.GET("/pegawai/:id", GetPegawaiByID) // pegawai/1 | pegawai/2

	e.Start(":8000")

}
