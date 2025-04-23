package models

import "time"

type Author struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthorInput используется для создания и обновления авторов
type AuthorInput struct {
	Name string `json:"name" binding:"required"`
}

// AuthorFilter используется для фильтрации авторов
type AuthorFilter struct {
	Name string `form:"name" json:"name"`
}
