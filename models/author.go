package models

type Author struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// AuthorInput используется для создания и обновления авторов
type AuthorInput struct {
	Name string `json:"name" binding:"required"`
}

// AuthorFilter используется для фильтрации авторов
type AuthorFilter struct {
	Name string `form:"name" json:"name"`
}
