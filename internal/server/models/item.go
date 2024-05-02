package models

import (
	"github.com/google/uuid"
)

type Item struct {
	ID     uuid.UUID `db:"id"`
	UserID uuid.UUID `db:"user_id"`
	Data   []byte    `db:"data"`
	Meta   []byte    `db:"meta"`
	Type   string    `db:"type"`
}
