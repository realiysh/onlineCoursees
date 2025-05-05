package models

type Course struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}

// структура для POST-запроса
type CourseInput struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}
