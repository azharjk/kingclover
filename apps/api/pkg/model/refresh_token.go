package model

import "github.com/google/uuid"

type RefreshToken struct {
	ID      uint
	Content *string
	UserID  uuid.UUID
}
