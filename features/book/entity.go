package book

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	Judul       string `validate:"min=3,omitempty"`
	TahunTerbit int    `validate:"numeric,gte=1900,lte=2050,omitempty"`
	Penulis     string `validate:"min=3,omitempty"`
	Pemilik     string
}

type BookHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	MyBook() echo.HandlerFunc
	GetAllBook() echo.HandlerFunc
}

type BookService interface {
	Add(token interface{}, newBook Core) (Core, error)
	Update(token interface{}, bookID int, updatedData Core) (Core, error)
	Delete(token interface{}, bookID int) error
	MyBook(token interface{}) ([]Core, error)
	GetAllBook() ([]Core, error)
}

type BookData interface {
	Add(userID int, newBook Core) (Core, error)
	Update(userID int, bookID int, updatedData Core) (Core, error)
	Delete(userID int, bookID int) error
	MyBook(userID int) ([]Core, error)
	GetAllBook() ([]Core, error)
}
