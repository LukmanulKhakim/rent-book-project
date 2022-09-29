package controller

import (
	"fmt"
	"project-rent/model"
)

type BookController struct {
	Model model.BookModel
}

func (bc BookController) GetAll() ([]model.Book, error) {
	res, err := bc.Model.GetAll()
	if err != nil {
		fmt.Println("Error on GetAll control", err.Error())
		return nil, err
	}
	return res, nil
}

func (bc BookController) GetBookId(ID_Book uint) (model.Book, error) {
	result, err := bc.Model.GetBookId(ID_Book)
	if err != nil {
		fmt.Println("Error on GetWhere", err.Error())
		return model.Book{}, err
	}
	return result, nil
}

func (bc BookController) Add(newBook model.Book) (model.Book, error) {
	res, err := bc.Model.Add(newBook)
	if err != nil {
		fmt.Println("Error on Add control", err.Error())
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookController) Edit(updatedBooks model.Book) (model.Book, error) {
	res, err := bc.Model.Edit(updatedBooks)
	if err != nil {
		fmt.Println("Error on Edit control", err.Error())
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookController) Delete(deletedBook model.Book) (model.Book, error) {
	res, err := bc.Model.Delete(deletedBook)
	if err != nil {
		fmt.Println("Error on Delete control", err.Error())
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookController) NotRent() ([]model.Book, error) {
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
