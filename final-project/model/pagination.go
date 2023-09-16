package model

type Pagination struct {
	Limit  *int   `form:"limit" binding:"omitempty,min=1"`
	Offset int    `form:"offset" binding:"min=0"`
	Search string `form:"search"`
}
