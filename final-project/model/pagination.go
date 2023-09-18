package model

type Paginator struct {
	Limit    int   `json:"limit"`
	Offset   int   `json:"offset"`
	Page     int   `json:"page"`
	LastPage int   `json:"last_page"`
	Total    int64 `json:"total"`
}

type ListQueryParam struct {
	Limit  *int   `form:"limit" binding:"omitempty,min=1"`
	Offset int    `form:"offset" binding:"min=0"`
	Search string `form:"search"`
}
