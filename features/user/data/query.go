package data

import (
	"api/features/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	res := User{}

	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return ToCore(res), nil
}
func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	cnv := CoreToData(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		return user.Core{}, err
	}

	newUser.ID = cnv.ID

	return newUser, nil
}
func (uq *userQuery) Profile(id uint) (user.Core, error) {
	res := User{}
	if err := uq.db.Where("id = ? AND deleted_at IS NULL", id).First(&res).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return user.Core{}, err
	}

	return ToCore(res), nil
}

func (uq *userQuery) Update(id uint, updateData user.Core) (user.Core, error) {
	dataModel := CoreToData(updateData)
	tx := uq.db.Model(User{}).Where("id = ?", id).Updates(&dataModel)
	if tx.Error != nil {
		log.Println("Update query error", tx.Error.Error())
		return user.Core{}, tx.Error
	}

	if tx.RowsAffected < 1 {
		log.Println("Rows affected update error")
		return user.Core{}, errors.New("user not found")
	}

	return ToCore(dataModel), nil
}

func (uq *userQuery) Deactive(id uint) error {
	tx := uq.db.Delete(User{}, id)
	if tx.Error != nil {
		log.Println("Delete query error", tx.Error.Error())
		return tx.Error
	}

	if tx.RowsAffected < 0 {
		log.Println("Rows affected delete error")
		return errors.New("user not found")
	}

	return nil
}
