package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id_user  int
	Nama     int
	Addres   string
	Email    string
	Password string
}

type UserModel struct {
	DB *gorm.DB
}

func (um UserModel) Register(NewUser User) (User, error) {
	Pass, err := bcrypt.GenerateFromPassword([]byte(NewUser.Password), bcrypt.MinCost)
	if err != nil {
		return NewUser, err
	}

	NewUser.Password = string(Pass)
	if err := um.DB.Save(&NewUser).Error; err != nil {
		return NewUser, err
	}
	return NewUser, nil
}

func (um UserModel) Login(Email, Password string) (User, error) {
	var user User
	var err error

	if err = um.DB.Where("Email = ?", Email).First(&user).Error; err != nil {
		fmt.Println("email wrong")
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if err != nil {
		fmt.Println("password wrong")
		return user, err
	}
	return user, nil
}

func (um UserModel) GetAll() ([]User, error) {
	var user []User
	err := um.DB.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (um UserModel) Update(UpdateUser User, Id_user int) (User, error) {
	var user User
	err := um.DB.Where("Id_user", UpdateUser.Id_user).Error
	if err != nil {
		return user, err
	}

	user.Nama = UpdateUser.Nama
	user.Addres = UpdateUser.Addres
	user.Password = UpdateUser.Password
	user.Email = UpdateUser.Email

	err = um.DB.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (um UserModel) NonAktive(Id_user int) (User, error) {

}
