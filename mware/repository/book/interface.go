package book

import "apiex/mware/entity"

type Book interface {
	Insert(newBook entity.Book) (entity.Book, error)
}
