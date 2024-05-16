package models

import (
	"github.com/google/uuid"
)

type Item struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Type   string
	Data   []byte
	Meta   []byte
}
