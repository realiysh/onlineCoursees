package models

type Course struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Title    string  `json:"title"`
	AuthorID uint    `json:"author_id"`
	Price    float64 `json:"price"`
}

// CourseInput используется для создания и обновления курсов
type CourseInput struct {
	Title    string  `json:"title" binding:"required"`
	AuthorID uint    `json:"author_id" binding:"required"`
	Price    float64 `json:"price" binding:"required,min=0"`
}

// CourseFilter используется для фильтрации курсов
type CourseFilter struct {
	Title    string  `form:"title" json:"title"`
	AuthorID uint    `form:"author_id" json:"author_id"`
	MinPrice float64 `form:"min_price" json:"min_price"`
	MaxPrice float64 `form:"max_price" json:"max_price"`
}
