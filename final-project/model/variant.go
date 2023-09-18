package model

import (
	"time"
)

type Variant struct {
	ID          uint      `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UUID        string    `json:"uuid" gorm:"not null"`
	VariantName string    `json:"variant_name" gorm:"not null"`
	Quantity    int       `json:"quantity" gorm:"not null"`
	ProductID   uint      `json:"-"`
	Product     *Product  `json:"product,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewVariant struct {
	VariantName string `form:"variant_name" binding:"required"`
	Quantity    *int   `form:"quantity" binding:"required,min=0"`
	ProductID   string `form:"product_id" binding:"required,uuid"`
}

type UUIDVariant struct {
	UUID string `uri:"uuid" binding:"required,uuid"`
}

type ExistingVariant struct {
	VariantName string `form:"variant_name" binding:"required"`
	Quantity    *int   `form:"quantity" binding:"required,min=0"`
	ProductID   string `form:"product_id" binding:"omitempty,uuid"`
}
