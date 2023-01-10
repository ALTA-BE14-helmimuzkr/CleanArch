package services

import (
	"api/features/user"
	"api/helper"
	"errors"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{qry: ud}
}

func (uuc *userUseCase) Login(email, password string) (string, user.Core, error) {
	res, err := uuc.qry.Login(email)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data user not found"
		} else {
			msg = "internal server error"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)); err != nil {
		// log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password doesnt match")
	}

	strToken := helper.GenerateToken(res.ID)

	return strToken, res, nil
}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {
	hashed, err := helper.HashPassword(newUser.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.Core{}, errors.New("internal server error")
	}

	newUser.Password = string(hashed)
	res, err := uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data user duplicate"
		} else {
			msg = "internal server error"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("data user not found")
	}
	res, err := uuc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data user not found"
		} else {
			msg = "internal server error"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (uuc *userUseCase) Update(token interface{}, updateData user.Core) (user.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("invalid user id")
	}

	res, err := uuc.qry.Update(uint(id), updateData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data user not found"
		} else {
			msg = "internal server error"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (uuc *userUseCase) Deactive(token interface{}) error {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return errors.New("id user not found")
	}

	err := uuc.qry.Deactive(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data user not found"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}
	return nil
}
