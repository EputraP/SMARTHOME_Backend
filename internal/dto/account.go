package dto

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username  string         `json:"username" gorm:"type:varchar;not null;unique"`
	LastLogin time.Time      `json:"last_login"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string `validate:"required, min=2, max=100" json:"username"`
	Password string `validate:"required, min=2, max=100" json:"password"`
}

