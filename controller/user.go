package controller

import (
	"fmt"
	"project-rent/model"
)

type UserController struct {
	Model model.UserModel
}

func (uc UserController) Register(NewReg model.User) (model.User, error) {

	NewUser, err := uc.Model.Register(NewReg)
	if err != nil {
		fmt.Println("Eror Register")
		return model.User{}, err
	}
	return NewUser, nil
}

func (uc UserController) Login(Email string, Password string) (model.User, error) {
	//var userRequest model.User
	user, err := uc.Model.Login(Email, Password)
	if err != nil {
		fmt.Println("eror login controll")
		return user, err
	}
	return user, nil
}

func (uc UserController) GetAll() ([]model.User, error) {
	res, err := uc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UserController) Edit(UpdateUser model.User) (model.User, error) {

	res, err := uc.Model.Edit(UpdateUser)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}
