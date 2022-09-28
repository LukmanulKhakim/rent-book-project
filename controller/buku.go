package controller

import (
	"fmt"
	"project-rent/model"
)

type BookControl struct {
	Model model.BookModel
}

func (bc BookControl) GetAll() ([]model.Book, error) {
	result, err := bc.Model.GetAll()
	if err != nil {
		fmt.Println("Error on GetAll", err.Error())
		return nil, err
	}
	return result, nil
}

func (bc BookControl) Add(newBook model.Book) (model.Book, error) {
	addBook, err := bc.Model.Add(newBook)
	if err != nil {
		fmt.Println("Error on Add", err.Error())
		return model.Book{}, err
	}
	return addBook, nil
}

func (bc BookControl) Edit(updatedBooks model.Book) (model.Book, error) {
	result, err := bc.Model.Edit(updatedBooks)
	if err != nil {
		fmt.Println("Error on Edit", err.Error())
		return model.Book{}, err
	}
	return result, nil
}

func (bc BookControl) Delete(deletedBook model.Book) (model.Book, error) {
	result, err := bc.Model.Delete(deletedBook)
	if err != nil {
		fmt.Println("Error on Delete", err.Error())
		return model.Book{}, err
	}
	return result, nil
}

// func (bc BooksControl) GetWhere(_title string) ([]model.Book, error) {
// 	result, err := bc.Model.Search(_title)
// 	if err != nil {
// 		fmt.Println("Error on GetWhere", err.Error())
// 		return nil, err
// 	}
// 	return result, nil
// }