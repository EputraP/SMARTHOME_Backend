package dto

import "github.com/google/uuid"

type RegisterBody struct {
	Password string `json:"password" binding:"required"`
	UserName string `json:"username" binding:"required"`
}

type RegisterResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}
