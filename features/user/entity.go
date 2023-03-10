package user

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Nama     string `validate:"omitempty,min=3"`
	Email    string `validate:"omitempty,email"`
	Alamat   string
	HP       string `validate:"omitempty,min=9,max=13"`
	Password string `validate:"omitempty,min=4"`
}

type UserHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Deactive() echo.HandlerFunc
}

type UserService interface {
	Login(email, password string) (string, Core, error)
	Register(newUser Core) (Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, updateData Core) (Core, error)
	Deactive(token interface{}) error
}

type UserData interface {
	Login(email string) (Core, error)
	Register(newUser Core) (Core, error)
	Profile(id uint) (Core, error)
	Update(id uint, updateData Core) (Core, error)
	Deactive(id uint) error
}
