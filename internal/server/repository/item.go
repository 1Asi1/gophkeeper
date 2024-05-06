package repository

import (
	"context"

	"github.com/google/uuid"
)

type Item struct {
	ID     uuid.UUID `db:"id"`
	UserID uuid.UUID `db:"user_id"`
	Type   string    `db:"type"`
	Data   []byte    `db:"data"`
	Meta   []byte    `db:"meta"`
}

type ItemRepository interface {
	CreateItem(context.Context, Item) (uuid.UUID, error)
	UpdateItem(context.Context, Item) error
	GetItem(context.Context, uuid.UUID) (Item, error)
	GetItems(context.Context, uuid.UUID) ([]Item, error)
	GetItemsByType(context.Context, uuid.UUID, string) ([]Item, error)
	DeleteItem(context.Context, uuid.UUID) error
}

func (s *Store) CreateItem(ctx context.Context, model Item) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (s *Store) UpdateItem(ctx context.Context, model Item) error {
	return nil
}

func (s *Store) GetItem(ctx context.Context, itemID uuid.UUID) (Item, error) {
	return Item{}, nil
}

func (s *Store) GetItems(ctx context.Context, userID uuid.UUID) ([]Item, error) {
	return []Item{}, nil
}

func (s *Store) GetItemsByType(ctx context.Context, userID uuid.UUID, itemType string) ([]Item, error) {
	return []Item{}, nil
}

func (s *Store) DeleteItem(ctx context.Context, itemID uuid.UUID) error {
	return nil
}
