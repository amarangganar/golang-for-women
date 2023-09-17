package model

import "time"

type Variant struct {
	ID          uint   `gorm:"primaryKey;AUTO_INCREMENT"`
	UUID        string `gorm:"not null"`
	VariantName string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	ProductID   uint
	Product     Product
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
