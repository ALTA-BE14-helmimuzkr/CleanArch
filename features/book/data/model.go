package data

import (
	"api/features/book"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Judul       string
	TahunTerbit int
	Penulis     string
	UserID      uint
}

type BookPemilik struct {
	ID          uint
	Judul       string
	TahunTerbit int
	Penulis     string
	Nama        string
}

func ToCore(data Books) book.Core {
	return book.Core{
		ID:          data.ID,
		Judul:       data.Judul,
		TahunTerbit: data.TahunTerbit,
		Penulis:     data.Penulis,
	}
}

func ToCoreSlice(data []BookPemilik) []book.Core {
	books := []book.Core{}
	for _, v := range data {
		book := book.Core{}
		book.ID = v.ID
		book.Judul = v.Judul
		book.Penulis = v.Penulis
		book.TahunTerbit = v.TahunTerbit
		book.Pemilik = v.Nama

		books = append(books, book)
	}

	return books
}

func CoreToData(data book.Core) Books {
	return Books{
		Model:       gorm.Model{ID: data.ID},
		Judul:       data.Judul,
		Penulis:     data.Penulis,
		TahunTerbit: data.TahunTerbit,
	}
}
