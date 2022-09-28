package controller

import (
	"fmt"
	"project-rent/model"
)

type RentController struct {
	Model model.RentModel
}

func (rc RentController) GetUserRent(ID_User uint) ([]model.Rent, error) {
	res, err := rc.Model.GetUserRent(ID_User)
	if err != nil {
		fmt.Println("Eror get user rent")
		return nil, err
	}
	return res, nil
}

func (rc RentController) AddRent(newRent model.Rent) (model.Rent, error) {
	res, err := rc.Model.AddRent(newRent)
	if err != nil {
		return model.Rent{}, err
	}
	return res, nil
}
