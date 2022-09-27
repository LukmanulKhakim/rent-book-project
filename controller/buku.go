package controller

import (
	"fmt"
	"log"
	"project-rent/model"
)

type BookController struct {
	Model model.BookModel
}

func (c BookController) GetAll() ([]model.Book, error) {
	temp, err := c.Model.GetAll()
	if err != nil {
		fmt.Println("Error on GetAll", err.Error())
		return nil, err
	}
	return temp, nil
}

func (c BookController) GetWhere(_title string) ([]model.Book, error) {
	temp, err := c.Model.GetWhere(_title)
	if err != nil {
		fmt.Println("Error on GetWhere", err.Error())
		return nil, err
	}
	return temp, nil
}

func (c BookController) Insert(newBook model.Book) (model.Book, error) {
	result, err := c.Model.Insert(newBook)
	if err != nil {
		log.Println("Wrong Insert", err.Error())
		return model.Book{}, err
	}
	return result, nil
}

func (c BookController) Edit(updatedBook model.Book) (model.Book, error) {
	result, err := c.Model.Edit(updatedBook)
	if err != nil {
		log.Println("Wrong Edit", err.Error())
		return model.Book{}, err
	}
	return result, nil
}

func (c BookController) Delete(deletedBook model.Book) (model.Book, error) {
	result, err := c.Model.Delete(deletedBook)
	if err != nil {
		log.Println("Wrong Delete", err.Error())
		return model.Book{}, err
	}
	return result, nil
}
