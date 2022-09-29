package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul      string
	Deskripsi  string
	Is_Rent    bool
	Is_Deleted bool
	ID_User    uint
	Rents      []Rent `gorm:"foreignKey:ID_Buku;"`
}

type BookModel struct {
	//DB *gorm.DB
	DB *gorm.DB
}

func (bm BookModel) GetAll() ([]Book, error) {
	var res []Book
	err := bm.DB.Find(&res).Error
	if err != nil {
		fmt.Println("Error on GetAll Model", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm BookModel) GetBookId(ID_Book uint) (Book, error) {
	var res Book
	err := bm.DB.First(&res, ID_Book).Error
	if err != nil {
		fmt.Println("Error ", err.Error())
		return Book{}, err
	}
	return res, nil
}

func (bm BookModel) Add(newBook Book) (Book, error) {
	err := bm.DB.Save(&newBook).Error
	if err != nil {
		fmt.Println("Error on Add Model", err.Error())
		return Book{}, err
	}
	return newBook, nil
}

func (bm BookModel) Edit(bukuEdit Book) (Book, error) {
	err := bm.DB.Save(&bukuEdit).Error
	if err != nil {
		fmt.Println("Error on Edit Model", err.Error())
		return Book{}, err
	}
	return bukuEdit, nil
}

func (bm BookModel) Delete(deletedBook Book) (Book, error) {
	err := bm.DB.Delete(&deletedBook).Error
	if err != nil {
		fmt.Println("Error on Delete Model", err.Error())
		return Book{}, err
	}
	return deletedBook, nil
}

func (bm BookModel) NotRent(ID_User uint) ([]Book, error) {
	var result []Book
	err := bm.DB.Where("is_Rent = ? AND ID_User != ?", 0, ID_User).Find(&result).Error
	if err != nil {
		fmt.Println("Error on NotRent", err.Error())
		return nil, err
	}
	return result, nil
}

func (bm BookModel) GetMyBook(ID_User uint) ([]Book, error) {
	var res []Book
	err := bm.DB.Where("ID_User =?", ID_User).Find(&res).Error
	if err != nil {
		fmt.Println("Error on GetAll Model", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm BookModel) Search(judul string) ([]Book, error) {
	var res []Book
	err := bm.DB.Where(&Book{Judul: judul}).Find(&res).Error
	if err != nil {
		fmt.Println("Error search book model", err.Error())
		return nil, err
	}
	return res, nil
}
