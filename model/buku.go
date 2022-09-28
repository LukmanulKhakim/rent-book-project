package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul      string `gorm:"size:255;not null"`
	Deskripsi  string `gorm:"size:255;not null"`
	Is_Rent    bool
	Is_Deleted bool
	ID_User    uint
	Rents      []Rent `gorm:"foreignKey:ID_Buku"`
	//Id_Category uint
}

type BookModel struct {
	//DB *gorm.DB
	DB *gorm.DB
}

func (bm BookModel) GetAll() ([]Book, error) {
	var res []Book
	err := bm.DB.Find(&res).Error
	if err != nil {
		fmt.Println("Error on Query", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm BookModel) Add(newBook Book) (Book, error) {
	err := bm.DB.Save(&newBook).Error
	if err != nil {
		fmt.Println("Error on Create", err.Error())
		return Book{}, err
	}
	return newBook, nil
}

func (bm BookModel) Edit(bukuEdit Book) (Book, error) {
	err := bm.DB.Save(&bukuEdit).Error
	if err != nil {
		fmt.Println("Error on Edit", err.Error())
		return Book{}, err
	}
	return bukuEdit, nil
}

func (bm BookModel) Delete(deletedBook Book) (Book, error) {
	err := bm.DB.Delete(&deletedBook).Error
	if err != nil {
		fmt.Println("Error on Delete", err.Error())
		return Book{}, err
	}
	return deletedBook, nil
}

// func (bm BookModel) Search(judul string) ([]Book, error) {
// 	var result []Book
// 	err := bm.DB.Where(&Book{Judul: judul}).Find(&result).Error
// 	if err != nil {
// 		fmt.Println("Error on Query", err.Error())
// 		return nil, err
// 	}
// 	return result, nil
// }
