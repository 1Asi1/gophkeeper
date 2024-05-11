package repository

import (
	"context"
	"errors"
	"fmt"

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
	id := uuid.New()
	model.ID = id

	query := `
		INSERT INTO tbl_items(id, user_id, type, data, meta)
		VALUES (:id, :user_id, :type, :data, :meta)`

	result, err := s.db.NamedExecContext(ctx, query, model)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("CreateItem: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("result.RowsAffected: %w", err)
	}

	if rows == 0 {
		return uuid.UUID{}, fmt.Errorf("rows empty %w", errors.New("no rows affected"))
	}

	return id, nil
}

func (s *Store) UpdateItem(ctx context.Context, model Item) error {
	query := `
		UPDATE
		    tbl_items
		SET
		    type  = :type,
		    data  = :data,
		    meta = :meta
		WHERE
		    id = :id`

	result, err := s.db.NamedExecContext(ctx, query, model)
	if err != nil {
		return fmt.Errorf("UpdateItem: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected error: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("result.RowsAffected error: %w", errors.New("no rows affected"))
	}

	return nil
}

func (s *Store) GetItem(ctx context.Context, itemID uuid.UUID) (Item, error) {
	query := `
	SELECT
	    type,
		data ,
		meata
	FROM tbl_items
	WHERE id = $1
`
	var model Item
	err := s.db.GetContext(ctx, &model, query, itemID)
	if err != nil {
		return Item{}, fmt.Errorf("GetItem: %w", err)
	}

	return model, nil
}

func (s *Store) GetItems(ctx context.Context, userID uuid.UUID) ([]Item, error) {
	var data []Item

	query := `
	SELECT
	    type,
		data ,
		meata
	FROM tbl_items
	WHERE id = $1
`

	if err := s.db.SelectContext(ctx, &data, query, userID); err != nil {
		return nil, fmt.Errorf("GetItems error: %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf(" GetItems error: %w", errors.New("no rows affected"))
	}

	result := make([]Item, len(data))

	for i, v := range data {
		result[i] = Item{
			Type: v.Type,
			Data: v.Data,
			Meta: v.Meta,
		}
	}

	return result, nil
}

func (s *Store) GetItemsByType(ctx context.Context, userID uuid.UUID, itemType string) ([]Item, error) {
	var data []Item

	query := `
	SELECT
	    type,
		data ,
		meata
	FROM tbl_items
	WHERE 
	    id = $1, 
	    type = $2
`

	if err := s.db.SelectContext(ctx, &data, query, userID, itemType); err != nil {
		return nil, fmt.Errorf("GetItems error: %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf(" GetItems error: %w", errors.New("no rows affected"))
	}

	result := make([]Item, len(data))

	for i, v := range data {
		result[i] = Item{
			Type: v.Type,
			Data: v.Data,
			Meta: v.Meta,
		}
	}

	return result, nil
}

func (s *Store) DeleteItem(ctx context.Context, itemID uuid.UUID) error {
	query := `
		DELETE FROM tbl_items
		WHERE id = $1`

	result, err := s.db.ExecContext(ctx, query, itemID)
	if err != nil {
		return fmt.Errorf("DeleteItem error: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected error: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("result.RowsAffected error: %w", errors.New("no rows affected"))
	}

	return nil
}
