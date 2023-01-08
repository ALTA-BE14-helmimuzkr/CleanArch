package data

import (
	"api/features/book"
	"errors"
	"log"

	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.BookData {
	return &bookData{
		db: db,
	}
}

func (bd *bookData) Add(userID int, newBook book.Core) (book.Core, error) {
	// Convert book.Core to model Books
	cnv := CoreToData(newBook)
	// Assign userID to cnv
	cnv.UserID = uint(userID)

	// DB Create
	err := bd.db.Create(&cnv).Error
	if err != nil {
		return book.Core{}, err
	}

	// return result converting cnv to book.Core
	return ToCore(cnv), nil
}

func (bd *bookData) Update(userID int, bookID int, updatedData book.Core) (book.Core, error) {
	// Convert book.Core to model Books
	cnv := CoreToData(updatedData)

	// DB Update(value)
	tx := bd.db.Where("id = ? && user_id = ?", bookID, userID).Updates(&cnv)
	if tx.Error != nil {
		log.Println("update book query error :", tx.Error)
		return book.Core{}, tx.Error
	}

	// Rows affected checking
	if tx.RowsAffected <= 0 {
		log.Println("update book query error : data not found")
		return book.Core{}, errors.New("not found")
	}

	// return result converting cnv to book.Core
	return ToCore(cnv), nil
}

func (bd *bookData) Delete(userID int, bookID int) error {
	// DB Delete(table or value)
	tx := bd.db.Where("id = ? AND user_id = ?", bookID, userID).Delete(&Books{})
	if tx.Error != nil {
		log.Println("delete book query error :", tx.Error)
		return tx.Error
	}

	// Rows affected checking
	if tx.RowsAffected <= 0 {
		log.Println("delete book query error : data not found")
		return errors.New("not found")
	}

	return nil
}

func (bd *bookData) MyBook(userID int) ([]book.Core, error) {
	// Representation of []BookPemilik, for contain all data result db Find
	myBooks := []BookPemilik{}
	// DB Raw with parameter Select Join
	// and combine with Find(destination)
	err := bd.db.Raw("SELECT books.id, books.judul, books.tahun_terbit, books.penulis, users.nama FROM books JOIN users ON users.id = books.user_id WHERE books.user_id = ? AND books.deleted_at IS NULL", userID).Find(&myBooks).Error
	if err != nil {
		return nil, err
	}

	if myBooks == nil {
		return nil, errors.New("books not found")
	}

	// Convert []BookPemilik model to []book.Core
	var bookCore = ToCoreSlice(myBooks)

	return bookCore, nil
}

func (bd *bookData) GetAllBook() ([]book.Core, error) {
	// Representation of []BookPemilik, for contain all data result db Find
	books := []BookPemilik{}
	// DB Raw with parameter Select Join
	// and combine with Find(destination)
	err := bd.db.Raw("SELECT books.id, books.judul, books.tahun_terbit, books.penulis, users.nama FROM books JOIN users ON users.id = books.user_id WHERE books.deleted_at IS NULL").Find(&books).Error
	if err != nil {
		return nil, err
	}

	if books == nil {
		return nil, errors.New("books not found")
	}

	// Convert []BookPemilik model to []book.Core
	var bookCore = ToCoreSlice(books)

	return bookCore, nil
}
