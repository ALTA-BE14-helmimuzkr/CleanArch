package services

import (
	"api/features/book"
	"api/helper"
	"errors"
	"strings"
)

type bookSrv struct {
	data book.BookData
}

func New(d book.BookData) book.BookService {
	return &bookSrv{data: d}
}

func (bs *bookSrv) Add(token interface{}, newBook book.Core) (book.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return book.Core{}, errors.New("user not found")
	}

	res, err := bs.data.Add(userID, newBook)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user not found"
		} else {
			msg = "internal server error"
		}
		return book.Core{}, errors.New(msg)
	}

	return res, nil

}
func (bs *bookSrv) Update(token interface{}, bookID int, updatedData book.Core) (book.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return book.Core{}, errors.New("user not found")
	}

	res, err := bs.data.Update(userID, bookID, updatedData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book or user not found"
		} else {
			msg = "internal server error"
		}
		return book.Core{}, errors.New(msg)
	}

	return res, nil
}

func (bs *bookSrv) Delete(token interface{}, bookID int) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("user not found")
	}

	err := bs.data.Delete(userID, bookID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book or user not found"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}

	return nil
}

func (bs *bookSrv) MyBook(token interface{}) ([]book.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return nil, errors.New("user not found")
	}

	res, err := bs.data.MyBook(userID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book or user not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}

	if len(res) < 1 {
		return nil, errors.New("books not found")
	}

	return res, nil
}

func (bs *bookSrv) GetAllBook() ([]book.Core, error) {
	res, err := bs.data.GetAllBook()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "book not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}

	if len(res) < 1 {
		return nil, errors.New("books not found")
	}

	return res, nil
}
