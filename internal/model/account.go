package model

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username  string         `json:"username" gorm:"type:varchar;not null;unique"`
	Password string         `json:"password" gorm:"type:varchar"`
	CreatedOn time.Time      `json:"created_on"`
	LastLogin time.Time      `json:"last_login"`
}

