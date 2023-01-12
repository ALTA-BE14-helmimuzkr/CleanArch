package services

import (
	"api/features/book"
	"api/helper"
	"api/mocks"
	"errors"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	repo := mocks.NewBookData(t) // Create new data mock
	v := validator.New()
	srv := New(repo, v)

	// Case: user menambahkan buku baru, lalu ketika sukses akan muncul buku yang sudah ditambahkan
	t.Run("Add successfully", func(t *testing.T) {
		input := book.Core{
			Judul:       "Naruto",
			Penulis:     "Masashi Kishimoto",
			TahunTerbit: 1999,
		}
		resData := book.Core{
			ID:          1,
			Judul:       "Naruto",
			Penulis:     "Masashi Kishimoto",
			TahunTerbit: 1999,
			Pemilik:     "helmi",
		}

		// Mockup repo using native testify mock
		// repo := &data.Mockrepo{Mock: &mock.Mock{}}
		// repo.Mock.On("Add", data.UserCollection[0].ID, data.InputCollection[0]).Return(data.RespCollection[0], nil).Once()

		// Mockup repo using Mockrey
		repo.Mock.On("Add", 1, input).Return(resData, nil).Once()

		// Create new service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Add(token, input)
		fmt.Println("================", err)

		// Test
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, actual.ID)
		assert.Equal(t, resData.Judul, actual.Judul)
		assert.Equal(t, resData.Penulis, actual.Penulis)
		assert.Equal(t, resData.TahunTerbit, actual.TahunTerbit)
		assert.Equal(t, resData.Pemilik, actual.Pemilik)
		repo.AssertExpectations(t)
	})

	t.Run("Addd error id not found", func(t *testing.T) {
		input := book.Core{
			Judul:       "Naruto",
			Penulis:     "Masashi Kishimoto",
			TahunTerbit: 1999,
		}

		// Program service
		token := jwt.New(jwt.SigningMethodHS256)
		actual, err := srv.Add(token, input)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "id user not found")
		assert.Empty(t, actual)
	})

	// Case: user menambahkan buku baru, tetapi input tidak sesuai aturan validasi
	t.Run("Add error input invalid", func(t *testing.T) {
		input := book.Core{
			Judul:       "n",  // min 3 character
			Penulis:     "ma", // min 3 character
			TahunTerbit: 1800, // less than 1900
		}

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Add(token, input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "validation input failed")
		assert.Empty(t, actual)
	})

	// Case: user menambahkan buku baru, tetapi terdapat masalah pada database
	t.Run("Add error internal server", func(t *testing.T) {
		input := book.Core{
			Judul:       "Naruto",
			Penulis:     "Masashi Kishimoto",
			TahunTerbit: 1999,
		}
		resData := book.Core{}

		// Programming input and return method add in query
		repo.Mock.On("Add", 1, input).Return(resData, errors.New("server error")).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Add(token, input)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})
}

func TestUpdateBook(t *testing.T) {
	input := book.Core{Judul: "One Piece"}
	resData := book.Core{
		ID:          1,
		Judul:       "One Piece",
		TahunTerbit: 1999,
		Penulis:     "Masashi Kishimoto",
		Pemilik:     "helmi",
	}

	repo := mocks.NewBookData(t)
	v := validator.New()
	srv := New(repo, v)

	// Case: user melakukan update
	t.Run("Update successfully", func(t *testing.T) {
		// Programming input and return repo
		repo.On("Update", 1, 1, input).Return(resData, nil).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.Nil(t, err)
		assert.Equal(t, resData.Judul, actual.Judul)
		assert.Equal(t, resData.ID, actual.ID)
		assert.Equal(t, resData.Pemilik, actual.Pemilik)

		repo.AssertExpectations(t)
	})

	// Case: user melakukan update tetapi token tidak valid
	t.Run("Update error user not found", func(t *testing.T) {
		// Program service
		// strToken := helper.GenerateToken(1)
		// token := helper.ValidateToken(strToken)
		token := jwt.New(jwt.SigningMethodHS256)
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "id user not found")
		assert.Empty(t, actual)
	})

	// Case: user melakukan update tetapi terjadi error karena input invalid
	t.Run("Update error invalid", func(t *testing.T) {
		input := book.Core{
			Judul:       "n",  // min 3 character
			Penulis:     "ma", // min 3 character
			TahunTerbit: 1800, // less than 1900
		}

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "validation input failed")
		assert.Empty(t, actual)
	})

	// Case: user melakukan update tetapi terjadi error karena buku yang ingin diupdate tidak ditemukan
	t.Run("Update error book not found", func(t *testing.T) {
		// Programming input and return repo
		repo.On("Update", 1, 1, input).Return(book.Core{}, errors.New("not found")).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "book or user not found")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})

	// Case: user melakukan update, tetapi terdapat masalah pada database
	t.Run("Update error internal server", func(t *testing.T) {
		// Programming input and return repo
		repo.On("Update", 1, 1, input).Return(book.Core{}, errors.New("internal server error")).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})
}

func TestDeleteBook(t *testing.T) {
	repo := mocks.NewBookData(t)
	v := validator.New()
	srv := New(repo, v)

	// Case: user menghapus buku dan berhasil
	t.Run("Delete Success", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(nil).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := srv.Delete(token, 1)

		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})

	// Case: user menghapus buku, lalu terjadi error karena token invalid
	t.Run("Delete error user not found", func(t *testing.T) {
		// Program service
		token := jwt.New(jwt.SigningMethodHS256)
		err := srv.Delete(token, 1)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "id user not found")
	})

	// Case: user menghapus buku, lalu terjadi error karena buku tidak ditemukan
	t.Run("Delete error id not found", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(errors.New("not found")).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := srv.Delete(token, 1)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "book or user not found")

		repo.AssertExpectations(t)
	})

	// Case: user menghapus buku, lalu terjadi masalah pada database
	t.Run("Delete error internal server", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(errors.New("internal server error")).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := srv.Delete(token, 1)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")

		repo.AssertExpectations(t)
	})
}

func TestMyBook(t *testing.T) {
	repo := mocks.NewBookData(t)
	v := validator.New()
	srv := New(repo, v)

	resData := []book.Core{
		{
			ID:          1,
			Judul:       "Naruto",
			Penulis:     "Masashi Kishimoto",
			TahunTerbit: 1999,
		},
		{
			ID:          2,
			Judul:       "Dragon ball",
			Penulis:     "Akira toriyama",
			TahunTerbit: 1998,
		},
	}

	// Case: pimilik ingin melihat list buku yang dimilikinya
	t.Run("MyBook list succesfully", func(t *testing.T) {
		// Programming input and return repo
		repo.On("MyBook", 1).Return(resData, nil).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.MyBook(token)

		// Test
		for i := range actual {
			assert.Nil(t, err)
			assert.Equal(t, resData[i].ID, actual[i].ID)
			assert.Equal(t, resData[i].Judul, actual[i].Judul)
		}
	})

	// Case: pimilik ingin melihat list buku yang dimilikinya, tetapi tokennya invalid
	t.Run("MyBook error user", func(t *testing.T) {
		// Program service
		token := jwt.New(jwt.SigningMethodHS256)
		actual, err := srv.MyBook(token)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "id user not found")
		assert.Empty(t, actual)
	})

	// Case: pimilik ingin melihat list buku yang dimilikinya, tetapi error tidak ditemukan pada database
	t.Run("MyBook error id book or user", func(t *testing.T) {
		// Programming input and return repo
		repo.On("MyBook", 1).Return(nil, errors.New("not found")).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.MyBook(token)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "book or user not found")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})

	// Case: pimilik ingin melihat list buku yang dimilikinya, tetapi terdapat masalah pada database
	t.Run("My book error internal server", func(t *testing.T) {
		// Programming input and return repo
		repo.On("MyBook", 1).Return(nil, errors.New("internal server error")).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.MyBook(token)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})

	// Case: pimilik ingin melihat list buku yang dimilikinya, tetapi buku tidak ada buku yang ditemukan
	t.Run("MyBook error books", func(t *testing.T) {
		// Programming input and return repo
		repo.On("MyBook", 1).Return([]book.Core{}, nil).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.MyBook(token)

		// Test
		assert.NotNil(t, err)
		assert.Error(t, err, "books not found")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})
}

func TestGetAllBook(t *testing.T) {
	repo := mocks.NewBookData(t)
	v := validator.New()
	srv := New(repo, v)

	resData := []book.Core{
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
			Penulis:     "Akira toriyama",
			TahunTerbit: 1998,
			Pemilik:     "helmi",
		},
		{
			ID:          3,
			Judul:       "One piece",
			Penulis:     "Oda sensei",
			TahunTerbit: 1998,
			Pemilik:     "muzakir",
		},
	}

	// Case: user ingin menampilkan semua buku yang terdaftar
	t.Run("Get all book succesfully", func(t *testing.T) {
		// Programming input and return repo
		repo.On("GetAllBook").Return(resData, nil).Once()

		// Program service
		actual, err := srv.GetAllBook()

		// Test
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ID, actual[0].ID)
		assert.Equal(t, resData[0].Judul, actual[0].Judul)
		assert.Equal(t, resData[0].Pemilik, actual[0].Pemilik)
		assert.Equal(t, resData[1].ID, actual[1].ID)
		assert.Equal(t, resData[1].Pemilik, actual[1].Pemilik)
		assert.Equal(t, resData[2].ID, actual[2].ID)
		assert.Equal(t, resData[2].Pemilik, actual[2].Pemilik)
	})

	// Case: user ingin melihat list buku yang, tetapi error tidak ditemukan pada database
	t.Run("Get all book error not found", func(t *testing.T) {
		// Programming input and return repo
		repo.On("GetAllBook").Return(nil, errors.New("not found")).Once()

		// Program service
		actual, err := srv.GetAllBook()

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "books not found")
		assert.Nil(t, actual)

	})

	// Case: user ingin melihat list buku yang, tetapi terdapat masalah pada database
	t.Run("Get all book error server", func(t *testing.T) {
		// Programming input and return repo
		repo.On("GetAllBook").Return(nil, errors.New("internal server error")).Once()

		// Program service
		actual, err := srv.GetAllBook()

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		assert.Nil(t, actual)

	})

	// Case: user ingin melihat list buku yang, tetapi buku tidak ada buku yang ditemukan
	t.Run("Get all book error not found", func(t *testing.T) {
		// Programming input and return repo
		repo.On("GetAllBook").Return(nil, nil).Once()

		// Program service
		actual, err := srv.GetAllBook()

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "books not found")
		assert.Nil(t, actual)

	})
}
