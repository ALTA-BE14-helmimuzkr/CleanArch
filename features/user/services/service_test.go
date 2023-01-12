package services

import (
	"api/features/user"
	"api/helper"
	"api/mocks"
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)
	v := validator.New()
	srv := New(repo, v)

	// Case: user melakukan pendaftaran akun baru
	t.Run("Register successfully", func(t *testing.T) {
		// Prgramming input and return repo
		hashed, _ := helper.HashPassword("helmi")
		resData := user.Core{
			ID:       1,
			Nama:     "helmi",
			Email:    "helmi@gmail.com",
			Password: hashed,
			Alamat:   "depok",
			HP:       "081280888",
		}
		repo.On("Register", mock.Anything).Return(resData, nil).Once()

		// Program service
		input := user.Core{
			Nama:     "helmi",
			Email:    "helmi@gmail.com",
			Password: "helmi",
			Alamat:   "depok",
			HP:       "081280888",
		}
		actual, err := srv.Register(input)

		// Test
		assert.Nil(t, err)
		// Jika actual password(hashed) dan input password itu di compare dan hasilnya true, maka tidak akan menghasilkan error
		errCompare := bcrypt.CompareHashAndPassword([]byte(actual.Password), []byte(input.Password))
		assert.NoError(t, errCompare)

		assert.Equal(t, actual.ID, resData.ID)
		repo.AssertExpectations(t)
	})

	// t.Run("Register error on bcrypt", func(t *testing.T) {})

	// Case: user melakukan register tetapi terjadi error karena input invalid
	t.Run("Register error invalid", func(t *testing.T) {
		input := user.Core{
			Nama: "h",         // min 3 character
			HP:   "081222222", // min 9 and max 13 character
		}

		// Program service
		actual, err := srv.Register(input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "validation input failed")
		assert.Empty(t, actual)
	})

	// Case: user melakukan pendaftaran akun baru, tetapi terdapat masalah karena data yang didaftarkan sudah  terdaftar di dalam database
	t.Run("Register error data duplicate", func(t *testing.T) {
		input := user.Core{
			Nama:     "helmi",
			Email:    "helmi@gmail.com",
			Password: "helmi",
			Alamat:   "depok",
			HP:       "081280888",
		}

		// Programming input and return repo
		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("duplicated")).Once()

		// Program service
		actual, err := srv.Register(input)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "data user duplicate")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})

	// Case: user melakukan pendaftaran akun baru, tetapi terdapat masalah di database
	t.Run("Register error internal server", func(t *testing.T) {
		input := user.Core{
			Nama:     "helmi",
			Email:    "helmi@gmail.com",
			Password: "helmi",
			Alamat:   "depok",
			HP:       "081280888",
		}

		// Programming input and return repo
		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("database error")).Once()

		// Program service
		actual, err := srv.Register(input)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		assert.Empty(t, actual)
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)
	v := validator.New()
	srv := New(repo, v)

	t.Run("Login successfully", func(t *testing.T) {
		// input dan respond untuk mock data
		inputEmail := "helmi@gmail.com"
		// res dari data akan mengembalik password yang sudah di hash
		hashed, _ := helper.HashPassword("helmi")
		resData := user.Core{ID: uint(1), Nama: "helmi", Email: "helmi@gmail.com", HP: "08123456", Password: hashed}

		repo.On("Login", inputEmail).Return(resData, nil).Once() // simulasi method login pada layer data

		token, res, err := srv.Login(inputEmail, "helmi")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Login error not found", func(t *testing.T) {
		inputEmail := "muzakir@gmail.com"
		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("not found")).Once()

		token, res, err := srv.Login(inputEmail, "helmi")

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data user not found")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Login error internal server", func(t *testing.T) {
		inputEmail := "muzakir@gmail.com"
		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("database error")).Once()

		token, res, err := srv.Login(inputEmail, "helmi")

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Login error password doesnt match", func(t *testing.T) {
		inputEmail := "helmi@gmail.com"
		hashed, _ := helper.HashPassword("helmi")
		resData := user.Core{ID: uint(1), Nama: "helmi", Email: "helmi@gmail.com", HP: "08123456", Password: hashed}
		repo.On("Login", inputEmail).Return(resData, nil).Once()

		token, res, err := srv.Login(inputEmail, "muzakir")

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password doesnt match")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)
	v := validator.New()
	srv := New(repo, v)

	t.Run("Profile successfully", func(t *testing.T) {
		resData := user.Core{ID: uint(1), Nama: "helmi", Email: "helmi@gmail.com", HP: "08123456"}
		repo.On("Profile", uint(1)).Return(resData, nil).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		token.Valid = true
		res, err := srv.Profile(token)

		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Nama, res.Nama)
		assert.Equal(t, resData.HP, res.HP)
		repo.AssertExpectations(t)
	})

	t.Run("Jwt not valid", func(t *testing.T) {
		token := jwt.New(jwt.SigningMethodHS256)
		res, err := srv.Profile(token)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data user not found")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("Profile error user not found", func(t *testing.T) {
		repo.On("Profile", uint(4)).Return(user.Core{}, errors.New("not found")).Once()

		strToken := helper.GenerateToken(4)
		token := helper.ValidateToken(strToken)
		token.Valid = true
		res, err := srv.Profile(token)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data user not found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Profile error internal server", func(t *testing.T) {
		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("database error")).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		token.Valid = true
		res, err := srv.Profile(token)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)
	v := validator.New()
	srv := New(repo, v)

	// Case: user mengganti nama
	t.Run("Update successfully", func(t *testing.T) {
		input := user.Core{Nama: "muzakir"}

		resData := user.Core{ID: uint(1), Nama: "helmi", Email: "helmi@gmail.com", Alamat: "depok", HP: "081280868869"}
		repo.On("Update", uint(1), input).Return(resData, nil).Once()

		tokenStr := helper.GenerateToken(1)
		token := helper.ValidateToken(tokenStr)
		res, err := srv.Update(token, input)

		assert.NoError(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	// Case: id user tidak valid atau tidak ditemukan
	t.Run("Update error invalid id", func(t *testing.T) {
		input := user.Core{Nama: "helmi", Email: "helmi@gmail.com", Alamat: "depok", HP: "081222222"}
		token := jwt.New(jwt.SigningMethodHS256)
		res, err := srv.Update(token, input)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "invalid user id")
		assert.Empty(t, res.Nama)
		assert.Equal(t, input.ID, res.ID)
		repo.AssertExpectations(t)
	})

	// Case: user melakukan update tetapi terjadi error karena input invalid
	t.Run("Update error invalid", func(t *testing.T) {
		input := user.Core{
			Nama: "h",         // min 3 character
			HP:   "081222222", // min 9 and max 13 character
		}

		// Program service
		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.Update(token, input)

		// Test
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "validation input failed")
		assert.Empty(t, actual)
	})

	// Case: user mengganti nama, tetapi id tidak ditemukan??
	t.Run("Update error data not found", func(t *testing.T) {
		input := user.Core{Nama: "muzakir"}

		resData := user.Core{}
		repo.On("Update", uint(1), input).Return(resData, errors.New("not found")).Once()

		tokenStr := helper.GenerateToken(uint(1))
		token := helper.ValidateToken(tokenStr)
		res, err := srv.Update(token, input)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "data user not found")
		assert.Empty(t, res.Nama)
		assert.NotEqual(t, input.Nama, res.Nama)
		repo.AssertExpectations(t)
	})

	// Case: user ingin mengupdate data, tetapi terdapat error pada database
	t.Run("Update error internal server", func(t *testing.T) {
		input := user.Core{Nama: "muzakir"}

		resData := user.Core{}
		repo.On("Update", uint(1), input).Return(resData, errors.New("database error")).Once()

		tokenStr := helper.GenerateToken(uint(1))
		token := helper.ValidateToken(tokenStr)
		res, err := srv.Update(token, input)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		assert.Empty(t, res.Nama)
		assert.NotEqual(t, input.Nama, res.Nama)
		repo.AssertExpectations(t)
	})
}

func TestDeactive(t *testing.T) {
	repo := mocks.NewUserData(t)
	v := validator.New()
	srv := New(repo, v)

	// Case: user melakukan deactive account dan berhasil tanpa mendapatkan error
	t.Run("Deactive succesfully", func(t *testing.T) {
		repo.On("Deactive", uint(1)).Return(nil).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := srv.Deactive(token)

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	// Case: user melakukan deactive account, tetapi token/id tidak valid
	t.Run("Deactive error id not found", func(t *testing.T) {
		token := jwt.New(jwt.SigningMethodHS256)
		err := srv.Deactive(token)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "id user not found")
	})

	// Case: user melakukan deactive account, tetapi id tidak ditemukan
	t.Run("Deactive error user not found", func(t *testing.T) {
		repo.On("Deactive", uint(1)).Return(errors.New("not found")).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := srv.Deactive(token)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "data user not found")
		repo.AssertExpectations(t)
	})

	// Case: user melakukan deactive account tetapi terjadi masalah pada database
	t.Run("Deactive error internal server error", func(t *testing.T) {
		repo.On("Deactive", uint(1)).Return(errors.New("database error")).Once()

		strToken := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		err := srv.Deactive(token)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		repo.AssertExpectations(t)
	})
}
