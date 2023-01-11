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
		assert.ErrorContains(t, err, "id user not found")
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
		assert.ErrorContains(t, err, "input new book invalid")
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
		assert.ErrorContains(t, err, "internal server error")
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

	t.Run("Update successfully", func(t *testing.T) {
		repo.On("Update", 1, 1, input).Return(resData, nil).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Update(token, 1, input)

		assert.Nil(t, err)
		assert.Equal(t, resData.Judul, actual.Judul)
		assert.Equal(t, resData.ID, actual.ID)
		assert.Equal(t, resData.Pemilik, actual.Pemilik)

		repo.AssertExpectations(t)
	})

	t.Run("Update error user not found", func(t *testing.T) {
		// Program service
		// strToken := helper.GenerateToken(1)
		// token := helper.ValidateToken(strToken)
		token := jwt.New(jwt.SigningMethodHS256)
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "id user not found")
		assert.Empty(t, actual)
	})

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
		assert.ErrorContains(t, err, "input update book invalid")
		assert.Empty(t, actual)
	})

	t.Run("Update error book not found", func(t *testing.T) {
		// Programming input and return repo
		repo.On("Update", 1, 1, input).Return(book.Core{}, errors.New("not found")).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "book or user not found")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})

	t.Run("Update error internal server", func(t *testing.T) {
		// Programming input and return repo
		repo.On("Update", 1, 1, input).Return(book.Core{}, errors.New("internal server error")).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Update(token, 1, input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})
}

func TestDeleteBook(t *testing.T) {
	repo := mocks.NewBookData(t)
	v := validator.New()
	srv := New(repo, v)
	t.Run("Delete Success", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(nil).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := srv.Delete(token, 1)

		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("Delete Error", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(errors.New("user id not found")).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := srv.Delete(token, 1)

		assert.NotNil(t, err)

		repo.AssertExpectations(t)
	})
}

func TestMyBook(t *testing.T) {
	repo := mocks.NewBookData(t)
	v := validator.New()
	srv := New(repo, v)

	// Case: user ingin melihat list buku yang dimilikinya
	t.Run("MyBook list succesfully", func(t *testing.T) {
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

		// Programming input and return repo
		repo.On("MyBook", 1).Return(resData, nil).Once()

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.MyBook(token)

		// Test
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ID, actual[0].ID)
		assert.Equal(t, resData[0].Judul, actual[0].Judul)
		assert.Equal(t, resData[1].ID, actual[1].ID)
		assert.Equal(t, resData[1].Judul, actual[1].Judul)
	})
}

func TestGetAllBook(t *testing.T) {
	repo := mocks.NewBookData(t)
	v := validator.New()
	srv := New(repo, v)

	// Case: user ingin menampilkan semua buku yang terdaftar
	t.Run("MyBook list succesfully", func(t *testing.T) {
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
}
