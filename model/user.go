package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	gorm.Model
	//Id_user  int `gorm:"primaryKey"`
	Nama     string
	Addres   string
	Email    string
	Password string                `gorm:"type:varchar(255)"`
	IsDel    soft_delete.DeletedAt `gorm:"softDelete:flag"`
	Rent     []Rent                `gorm:"foreignKey:ID_User"`
	Book     []Book                `gorm:"foreignKey:ID_User"`
}

type UserModel struct {
	DB *gorm.DB
}

// register
func (um UserModel) Register(NewUser User) (User, error) {
	Pass, err := bcrypt.GenerateFromPassword([]byte(NewUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return NewUser, err
	}

	NewUser.Password = string(Pass)
	if err := um.DB.Save(&NewUser).Error; err != nil {
		fmt.Println("Eror Regist")
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
	//fmt.Println(user, Email, Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if err != nil {
		fmt.Println("password wrong")
		return user, err
	}
	if err := um.DB.Save(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// tampilkan semua data
func (um UserModel) GetAll() ([]User, error) {
	var user []User
	err := um.DB.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (um UserModel) Edit(UpdateUser User) (User, error) {
	//var user User
	err := um.DB.Save(&UpdateUser).Error
	if err != nil {
		fmt.Println("Error edit", err.Error())
		return User{}, err
	}
	return UpdateUser, nil
}

// func (um UserModel) NonAktive(nonAktive User, IsDel int) (User, error) {
// 	var user User

// 	err := um.DB.Where("IsDel", 1).Update().Error
// 	if err != nil {
// 		fmt.Println("eror non Aktive")
// 		return user, err
// 	}
// 	return user, nil
// }

// func (um UserModel) Edit(UpdateUser User) (User, error) {
// 	var user User
// 	err := um.DB.Where("id=?", UpdateUser.Id_user).Updates(&user).Error
// 	if err != nil {
// 		return user, err
// 	}

// 	err = um.DB.Save(&user).Error
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil

// }
