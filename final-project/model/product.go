package model

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
