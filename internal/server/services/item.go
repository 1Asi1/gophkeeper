package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gophkeeper/internal/server/models"
	"gophkeeper/internal/server/repository"
)

type ItemService interface {
	Create(context.Context, models.Item) (uuid.UUID, error)
	Update(context.Context, models.Item) error
	Get(context.Context, uuid.UUID) (models.Item, error)
	GetAll(context.Context, uuid.UUID) ([]models.Item, error)
	GetAllByType(context.Context, uuid.UUID, string) ([]models.Item, error)
	Delete(context.Context, uuid.UUID) error
}

type Item struct {
	log     zerolog.Logger
	storage repository.Store
	ctx     context.Context
}

func NewItemService() *Item {
	return &Item{}
}

func (s *Item) Create(ctx context.Context, req models.Item) (uuid.UUID, error) {
	id, err := s.storage.CreateItem(ctx, repository.Item{
		ID:     req.ID,
		UserID: req.UserID,
		Type:   req.Type,
		Data:   req.Data,
		Meta:   req.Meta,
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("s.storage.CreateItem: %w", err)
	}

	return id, nil
}

func (s *Item) Update(ctx context.Context, req models.Item) error {
	err := s.storage.UpdateItem(ctx, repository.Item{
		ID:     req.ID,
		UserID: req.UserID,
		Type:   req.Type,
		Data:   req.Data,
		Meta:   req.Meta,
	})
	if err != nil {
		return fmt.Errorf("s.storage.UpdateItem: %w", err)
	}

	return nil
}

func (s *Item) Get(ctx context.Context, itemID uuid.UUID) (models.Item, error) {
	res, err := s.storage.GetItem(ctx, itemID)
	if err != nil {
		return models.Item{}, fmt.Errorf("s.storage.GetItem: %w", err)
	}

	return models.Item{
		ID:     res.ID,
		UserID: res.UserID,
		Type:   res.Type,
		Data:   res.Data,
		Meta:   res.Meta,
	}, nil
}

func (s *Item) GetAll(ctx context.Context, userID uuid.UUID) ([]models.Item, error) {
	data, err := s.storage.GetItems(ctx, userID)
	if err != nil {
		return []models.Item{}, fmt.Errorf("s.storage.GetItems: %w", err)
	}

	res := make([]models.Item, len(data))
	for i, v := range data {
		res[i] = models.Item{
			ID:     v.ID,
			UserID: v.UserID,
			Type:   v.Type,
			Data:   v.Data,
			Meta:   v.Meta,
		}
	}

	return res, nil
}

func (s *Item) GetAllByType(ctx context.Context, userID uuid.UUID, itemType string) ([]models.Item, error) {
	data, err := s.storage.GetItemsByType(ctx, userID, itemType)
	if err != nil {
		return []models.Item{}, fmt.Errorf("s.storage.GetItemsByType: %w", err)
	}

	res := make([]models.Item, len(data))
	for i, v := range data {
		res[i] = models.Item{
			ID:     v.ID,
			UserID: v.UserID,
			Type:   v.Type,
			Data:   v.Data,
			Meta:   v.Meta,
		}
	}

	return res, nil
}

func (s *Item) Delete(ctx context.Context, itemID uuid.UUID) error {
	err := s.storage.DeleteItem(ctx, itemID)
	if err != nil {
		return fmt.Errorf("s.storage.DeleteItem: %w", err)
	}

	return nil
}
