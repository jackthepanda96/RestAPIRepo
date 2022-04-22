package book

import (
	"apiex/mware/entity"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type BookRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *BookRepo {
	return &BookRepo{
		Db: db,
	}
}

func (br *BookRepo) Insert(newBook entity.Book) (entity.Book, error) {
	if err := br.Db.Create(&newBook).Error; err != nil {
		log.Warn(err)
		return entity.Book{}, errors.New("tidak bisa insert data")
	}
	return newBook, nil
}
