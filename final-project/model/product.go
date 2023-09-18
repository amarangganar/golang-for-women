package model

import (
	"time"
)

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UUID      string    `json:"uuid" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"not null"`
	AdminID   uint      `json:"-"`
	Admin     Admin     `json:"admin,omitempty"`
	Variants  []Variant `json:"variants,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewProduct struct {
	Name    string `form:"name" validate:"required"`
	File    string `form:"file" validate:"required"`
	AdminID *uint  `validate:"-"`
}

type UUIDProduct struct {
	UUID string `uri:"uuid" binding:"required,uuid"`
}

type ExistingProduct struct {
	Name string `form:"name" validate:"required"`
	File string `form:"file"`
}
