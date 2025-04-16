package models

import "github.com/golang-jwt/jwt/v5"

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
