package models

// PaginationParams содержит параметры пагинации
type PaginationParams struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}
