package models

type Course struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}
