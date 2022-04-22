package controller

import (
	"apiex/mware/delivery/view"
	"apiex/mware/delivery/view/pegawai"
	"apiex/mware/entity"
	pegawaiRepo "apiex/mware/repository/pegawai"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PegawaiController struct {
	Repo  pegawaiRepo.Pegawai
	Valid *validator.Validate
}

func New(repo pegawaiRepo.Pegawai, valid *validator.Validate) *PegawaiController {
	return &PegawaiController{
		Repo:  repo,
		Valid: valid,
	}
}

func (pc *PegawaiController) Insert(c echo.Context) error {
	var tmpPegawai pegawai.InsertPegawaiRequest

	if err := c.Bind(&tmpPegawai); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, pegawai.BadRequest())
	}

	if err := pc.Valid.Struct(tmpPegawai); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, pegawai.BadRequest())
	}

	newPegawai := entity.Pegawai{Nama: tmpPegawai.Nama, HP: tmpPegawai.HP, Password: tmpPegawai.Password}
	res, err := pc.Repo.Insert(newPegawai)

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil insert")
	return c.JSON(http.StatusCreated, pegawai.SuccessInsert(res))
}

func (pc *PegawaiController) GetAllPegawai(c echo.Context) error {
	id := ExtractTokenUserId(c)

	fmt.Println(id)

	res, err := pc.Repo.GetAll()

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil get all data")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil get all data",
		"status":  true,
		"data":    res,
	})
}

func (pc *PegawaiController) Login(c echo.Context) error {
	param := pegawai.LoginRequest{}

	if err := c.Bind(&param); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, pegawai.BadRequest())
	}

	if err := pc.Valid.Struct(param); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, pegawai.BadRequest())
	}

	data, err := pc.Repo.Login(param.HP, param.Password)

	if err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusNotFound, "HP atau Password tidak ditemukan")
	}

	res := pegawai.LoginResponse{Data: data}

	if res.Token == "" {
		token, _ := CreateToken(int(data.ID))
		res.Token = token
		return c.JSON(http.StatusOK, view.OK(res, "Berhasil login"))
	}

	return c.JSON(http.StatusOK, view.OK(res, "Berhasil login"))
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["expired"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("RH$SI4"))
}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return userId
	}
	return 0
}
