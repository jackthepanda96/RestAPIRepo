package book

import (
	bookView "apiex/mware/delivery/view/book"
	"apiex/mware/entity"
	"apiex/mware/repository/book"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type BookController struct {
	Repo  book.Book
	Valid *validator.Validate
}

func New(repo book.Book, valid *validator.Validate) *BookController {
	return &BookController{
		Repo:  repo,
		Valid: valid,
	}
}

func (pc *BookController) Insert(c echo.Context) error {
	var tmpPegawai bookView.BookInsert

	idOwner := ExtractTokenUserId(c)

	if err := c.Bind(&tmpPegawai); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, "fail")
	}

	if err := pc.Valid.Struct(tmpPegawai); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, "fail")
	}

	res, err := pc.Repo.Insert(entity.Book{Judul: tmpPegawai.Judul, Author: tmpPegawai.Author, Owner: uint(int(idOwner))})

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, "fail")
	}
	log.Info("berhasil insert")
	return c.JSON(http.StatusCreated, res)
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
