package models

type Course struct {
	ID       uint    `json:"id"`
	Title    string  `json:"title"`
	AuthorID uint    `json:"author_id"`
	Price    float64 `json:"price"`
}
