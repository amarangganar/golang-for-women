package model

import (
	"final_project/helpers"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint      `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UUID      string    `json:"uuid" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"-" gorm:"not null"`
	Products  []Product `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	admin.Password = helpers.HashPassword(admin.Password)

	return
}

type AdminSignIn struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,alphanum,min=8"`
}

type AdminRegister struct {
	Name string `form:"name" json:"name" binding:"required"`
	AdminSignIn
}
