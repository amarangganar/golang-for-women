package model

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
