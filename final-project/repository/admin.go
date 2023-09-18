package repository

import (
	"final_project/model"

	"github.com/google/uuid"
)

func (db *Database) SignIn(body *model.AdminSignIn) (*model.Admin, error) {
	admin := model.Admin{}

	err := db.DB.Where("email = ?", body.Email).First(&admin).Error

	return &admin, err
}

func (db *Database) Register(body *model.AdminRegister) (*model.Admin, error) {
	admin := model.Admin{
		UUID:     uuid.NewString(),
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	err := db.DB.Create(&admin).Error

	return &admin, err
}
