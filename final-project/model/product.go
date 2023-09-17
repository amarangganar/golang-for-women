package model

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey;AUTO_INCREMENT"`
	UUID      string `gorm:"not null"`
	Name      string `gorm:"not null"`
	ImageURL  string `gorm:"not null"`
	AdminID   uint
	Admin     Admin
	Variants  []Variant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewProduct struct {
	Name string `form:"name" validate:"required"`
	File string `form:"file" validate:"required"`
}

type UUIDProduct struct {
	UUID string `uri:"uuid" binding:"required,uuid"`
}

type ExistingProduct struct {
	UUID string `uri:"uuid" validate:"required,uuid"`
	Name string `form:"name" validate:"required"`
	File string `form:"file"`
}
