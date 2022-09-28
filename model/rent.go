package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	//IsReturn       bool
	ID_Buku        uint
	ID_User        uint
	Books_id       uint
	Books_IsRent   bool
	Books_Nama     string
	Books_Email    string
	Books_Addres   string
	Judul_Book     string
	Deskripsi_Book string
	Return_book    time.Time
}

type RentModel struct {
	DB *gorm.DB
}

func (rm RentModel) GetUserRent(UserID uint) ([]Rent, error) {
	var res []Rent
	err := rm.DB.Where("Books_IsRent=? AND ").Find(&res).Error
	if err != nil {
		fmt.Println("Eror Get User Rent", err.Error())
		return nil, err
	}
	return res, nil
}

func (rm RentModel) AddRent(newRent Rent) (Rent, error) {
	err := rm.DB.Save(&newRent).Error
	if err != nil {
		fmt.Println("Error on Create", err.Error())
		return Rent{}, err
	}
	return newRent, nil
}

func (rm RentModel) RetRent(retNow Rent) (Rent, error) {
	err := rm.DB.Save(retNow).Error
	if err != nil {
		fmt.Println("Error on Create", err.Error())
		return Rent{}, err
	}
	return retNow, nil
}
