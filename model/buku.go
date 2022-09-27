package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Buku    string
	Status  string
	Pinjam  bool
	Hapus   bool
	Id_User uint
}

type BookModel struct {
	db *gorm.DB
}

// type BookModel interface {
// 	GetAll() ([]Book, error)
// 	GetWhere(bookId int) (Book, error)
// 	Insert(Book) (Book, error)
// 	Edit(book Book, BookId int) (Book, error)
// 	Delete(bookId int) (Book, error)
// }

func (b BookModel) GetAll() ([]Book, error) {
	var temp []Book
	err := b.db.Find(&temp).Error
	if err != nil {
		fmt.Println("wrong query", err.Error())
		return nil, err
	}
	return temp, nil
}

func (b BookModel) GetWhere(_buku string) ([]Book, error) {
	var temp []Book
	err := b.db.Where(&Book{Buku: _buku}).Find(&temp).Error
	if err != nil {
		fmt.Println("Wrong Query", err.Error())
		return nil, err
	}
	return temp, nil

}

func (b BookModel) Insert(newBook Book) (Book, error) {
	err := b.db.Save(&newBook).Error
	if err != nil {
		fmt.Println("Wrong Insert", err.Error())
		return Book{}, err
	}
	return newBook, nil
}

func (b BookModel) Edit(updateBook Book) (Book, error) {
	err := b.db.Save(&updateBook).Error
	if err != nil {
		fmt.Println("Wrong Edit", err.Error())
		return Book{}, err
	}
	return updateBook, nil
}

func (b BookModel) Delete(deleteBook Book) (Book, error) {
	err := b.db.Delete(&deleteBook).Error
	if err != nil {
		fmt.Println("Wrong Edit", err.Error())
		return Book{}, err
	}
	return deleteBook, nil
}
