package handler

import (
	"api/features/user"
)

type UserReponse struct {
	ID     uint   `json:"id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
	HP     string `json:"hp"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:     data.ID,
		Nama:   data.Nama,
		Email:  data.Email,
		Alamat: data.Alamat,
		HP:     data.HP,
	}
}
