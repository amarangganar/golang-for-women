package model

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey;AUTO_INCREMENT"`
	UUID      string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	Products  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminSignIn struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,alphanum,min=8"`
}

type AdminRegister struct {
	Name string `form:"name" json:"name" binding:"required"`
	AdminSignIn
}
