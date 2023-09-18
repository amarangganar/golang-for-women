package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UUID      string    `json:"uuid" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"not null"`
	AdminID   uint      `json:"-"`
	Admin     Admin     `json:"admin"`
	Variants  []Variant `json:"variants" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.UUID = uuid.NewString()

	return
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
	UUID string `uri:"uuid" validate:"required,uuid"`
	Name string `form:"name" validate:"required"`
	File string `form:"file"`
}
