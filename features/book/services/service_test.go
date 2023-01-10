package services

import (
	"api/features/book"
	"api/helper"
	"api/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	bookData := mocks.NewBookData(t) // Create new data mock
	bookService := New(bookData)

	t.Run("Add Success", func(t *testing.T) {
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

		// Mockup bookData using native testify mock
		// bookData := &data.MockBookData{Mock: &mock.Mock{}}
		// bookData.Mock.On("Add", data.UserCollection[0].ID, data.InputCollection[0]).Return(data.RespCollection[0], nil).Once()

		// Mockup bookData using Mockrey
		bookData.Mock.On("Add", 1, input).Return(resData, nil).Once()

		// Create new service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := bookService.Add(token, input)

		// Test
		assert.Nil(t, err)
		assert.Equal(t, resData.Judul, actual.Judul)
		assert.Equal(t, resData.ID, actual.ID)
		assert.Equal(t, resData.Pemilik, actual.Pemilik)

		bookData.AssertExpectations(t)
	})
}

func TestUpdateBook(t *testing.T) {
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

	bookData := mocks.NewBookData(t)
	bookService := New(bookData)

	t.Run("Update Success", func(t *testing.T) {

		bookData.On("Update", 1, 1, input).Return(resData, nil).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := bookService.Update(token, 1, input)

		assert.Nil(t, err)
		assert.Equal(t, resData.Judul, actual.Judul)
		assert.Equal(t, resData.ID, actual.ID)
		assert.Equal(t, resData.Pemilik, actual.Pemilik)

		bookData.AssertExpectations(t)
	})
}

func TestDeleteBook(t *testing.T) {
	bookData := mocks.NewBookData(t)
	bookService := New(bookData)

	t.Run("Delete Success", func(t *testing.T) {
		bookData.On("Delete", 1, 1).Return(nil).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := bookService.Delete(token, 1)

		assert.Nil(t, err)

		bookData.AssertExpectations(t)
	})

	t.Run("Delete Error", func(t *testing.T) {
		bookData.On("Delete", 1, 1).Return(errors.New("user id not found")).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := bookService.Delete(token, 1)

		assert.NotNil(t, err)

		bookData.AssertExpectations(t)
	})
}
