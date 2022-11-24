package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccessToken struct {
	ID      uint
	Content *string
	UserID  uuid.UUID
	Expires time.Time
}

func (token *AccessToken) BeforeCreate(tx *gorm.DB) error {
	token.Expires = time.Now().Add(time.Minute * 1)
	return nil
}
