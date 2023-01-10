package data

import (
	"api/features/book"

	"github.com/stretchr/testify/mock"
)

type MockBookData struct {
	Mock *mock.Mock
}

var UserCollection = []struct {
	ID   int
	Nama string
}{
	{
		ID:   1,
		Nama: "helmi",
	},
	{
		ID:   2,
		Nama: "muzakir",
	},
}
var ServiceCollection = []book.Core{
	{
		Judul:       "Naruto",
		Penulis:     "Masashi Kishimoto",
		TahunTerbit: 1999,
	},
	{
		Judul:       "Dragon ball",
		Penulis:     "Akira Toriyama",
		TahunTerbit: 1998,
	},
}
var DataCollection = []book.Core{
	{
		Judul:       "Naruto",
		Penulis:     "Masashi Kishimoto",
		TahunTerbit: 1999,
		Pemilik:     "helmi",
	},
	{
		Judul:       "Dragon ball",
		Penulis:     "Akira Toriyama",
		TahunTerbit: 1998,
		Pemilik:     "muzakir",
	},
}

var RespCollection = []book.Core{
	{
		ID:          1,
		Judul:       "Naruto",
		Penulis:     "Masashi Kishimoto",
		TahunTerbit: 1999,
		Pemilik:     "helmi",
	},
	{
		ID:          2,
		Judul:       "Dragon ball",
		Penulis:     "Akira Toriyama",
		TahunTerbit: 1998,
		Pemilik:     "muzakir",
	},
}

func (m *MockBookData) Add(userID int, newBook book.Core) (book.Core, error) {
	args := m.Mock.Called(userID, newBook)
	return args.Get(0).(book.Core), nil
}

func (m *MockBookData) Update(userID int, bookID int, updatedData book.Core) (book.Core, error) {
	return book.Core{}, nil
}

func (m *MockBookData) Delete(userID int, bookID int) error {
	return nil
}

func (m *MockBookData) MyBook(userID int) ([]book.Core, error) {
	return []book.Core{}, nil
}

func (m *MockBookData) GetAllBook() ([]book.Core, error) {
	return []book.Core{}, nil
}
