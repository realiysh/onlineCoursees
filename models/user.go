package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password,omitempty"`
}

type JWTClaim struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PaginationParams используется для получения параметров пагинации
type PaginationParams struct {
	Page     int `form:"page" json:"page" binding:"min=1"`
	PageSize int `form:"limit" json:"limit" binding:"min=1,max=100"`
}

// UserFilter используется для фильтрации пользователей
type UserFilter struct {
	Username string `form:"username" json:"username"`
}

// UpdateUserInput используется для обновления данных пользователя
type UpdateUserInput struct {
	Username string `json:"username"`
}
