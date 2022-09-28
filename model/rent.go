package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	IsReturn    bool
	ID_Buku     uint
	ID_User     uint
	Return_book time.Time
}

type RentModel struct {
	DB *gorm.DB
}

func (rm RentModel) GetUserRent(UserID uint) ([]Book, error) {
	var res []Book
	err := rm.DB.Where(&Book{ID_User: UserID}).Find(&res).Error
	if err != nil {
		fmt.Println("Eror Get User Rent")
		return nil, err
	}
	return res, nil
}
