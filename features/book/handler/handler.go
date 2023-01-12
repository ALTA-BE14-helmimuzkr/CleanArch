package handler

import (
	"api/features/book"
	"api/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type bookHandle struct {
	srv book.BookService
}

func New(bs book.BookService) book.BookHandler {
	return &bookHandle{
		srv: bs,
	}
}

func (bh *bookHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdateBookRequest{}
		if err := c.Bind(&input); err != nil {
			log.Println(err)
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		cnv := ToCore(input)

		res, err := bh.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println(err)
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		book := ToResponse("add", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan buku", book))
	}
}

func (bh *bookHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdateBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		cnv := ToCore(input)

		bookID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		res, err := bh.srv.Update(c.Get("user"), bookID, *cnv)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		book := ToResponse("update", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses mengubah buku", book))
	}
}

func (bh *bookHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		err = bh.srv.Delete(c.Get("user"), bookID)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses menghapus buku"))
	}
}

func (bh *bookHandle) MyBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bh.srv.MyBook(c.Get("user"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		listRes := ToListResponse(res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses menampilkan koleksi buku", listRes))
	}
}

func (bh *bookHandle) GetAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bh.srv.GetAllBook()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		listRes := ToListResponse(res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses menampilkan semua buku", listRes))
	}
}
