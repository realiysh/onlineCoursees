package models

type Category struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name" binding:"required" gorm:"unique"`
	Description string   `json:"description"`
	Courses     []Course `json:"courses,omitempty" gorm:"foreignKey:CategoryID"`
}

// CategoryInput используется для создания и обновления категорий
type CategoryInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CategoryFilter используется для фильтрации категорий
type CategoryFilter struct {
	Name string `form:"name" json:"name"`
}
