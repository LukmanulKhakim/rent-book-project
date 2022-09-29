package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Return_book  time.Time
	Books_IsRent bool
	//Books_id       uint //id pemilik buku
	Books_Nama     string
	Books_Email    string
	Judul_Book     string
	Deskripsi_Book string
	ID_User        uint //id user pemimjam
	ID_Buku        uint //id buku

}

type RentModel struct {
	DB *gorm.DB
}

func (rm RentModel) GetUserRent(UserID uint) ([]Rent, error) {
	var res []Rent
	err := rm.DB.Where("Books_Is_Rent=? AND id_user = ?", 0, UserID).Find(&res).Error
	if err != nil {
		fmt.Println("Eror Get User Rent", err.Error())
		return nil, err
	}
	return res, nil
}

func (rm RentModel) AddRent(newRent Rent) (Rent, error) {
	err := rm.DB.Save(&newRent).Error
	if err != nil {
		fmt.Println("Error Add rent", err.Error())
		return Rent{}, err
	}
	return newRent, nil
}

func (rm RentModel) RetRent(retNow Rent) (Rent, error) {
	err := rm.DB.Save(retNow).Error
	if err != nil {
		fmt.Println("Error ret rent", err.Error())
		return Rent{}, err
	}
	return retNow, nil
}
