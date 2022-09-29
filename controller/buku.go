package controller

import (
	"fmt"
	"project-rent/model"
)

type BookControl struct {
	Model model.BookModel
}

func (bc BookControl) GetAll() ([]model.Book, error) {
	res, err := bc.Model.GetAll()
	if err != nil {
		fmt.Println("Error on GetAll control", err.Error())
		return nil, err
	}
	return res, nil
}

func (bc BookControl) Add(newBook model.Book) (model.Book, error) {
	res, err := bc.Model.Add(newBook)
	if err != nil {
		fmt.Println("Error on Add control", err.Error())
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookControl) Edit(updatedBooks model.Book) (model.Book, error) {
	res, err := bc.Model.Edit(updatedBooks)
	if err != nil {
		fmt.Println("Error on Edit control", err.Error())
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookControl) Delete(deletedBook model.Book) (model.Book, error) {
	res, err := bc.Model.Delete(deletedBook)
	if err != nil {
		fmt.Println("Error on Delete control", err.Error())
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookControl) NotRent() ([]model.Book, error) {
	res, err := bc.Model.NotRent()
	if err != nil {
		fmt.Println("Error on NotRent control", err.Error())
		return nil, err
	}
	return res, nil
}

// func (bc BooksControl) GetWhere(_title string) ([]model.Book, error) {
// 	result, err := bc.Model.Search(_title)
// 	if err != nil {
// 		fmt.Println("Error on GetWhere", err.Error())
// 		return nil, err
// 	}
// 	return result, nil
// }
