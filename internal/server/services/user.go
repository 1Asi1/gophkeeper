package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"gophkeeper/internal/server/models"
	"gophkeeper/internal/server/repository"
)

type IUser interface {
	Create(context.Context, models.User) (uuid.UUID, error)
	Get(context.Context, string) (models.User, error)
	ComparePassword(models.User, string) (bool, error)
}

type User struct {
	log     zerolog.Logger
	storage repository.Store
	ctx     context.Context
}

func NewUserService() *User {
	return &User{}
}

func (s *User) Create(ctx context.Context, req models.User) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}

	user := repository.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	id, err := s.storage.Create(ctx, user)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("s.storage.Create: %w", err)
	}

	return id, nil
}

func (s *User) Get(ctx context.Context, email string) (models.User, error) {
	res, err := s.storage.Get(ctx, email)
	if err != nil {
		return models.User{}, fmt.Errorf("s.storage.Get: %w", err)
	}

	return models.User{
		ID:       res.ID.String(),
		Email:    res.Email,
		Password: string(res.Password),
	}, err
}

func (s *User) ComparePassword(req models.User, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(password)); err != nil {
		return false, fmt.Errorf(" bcrypt.CompareHashAndPassword: %w", err)
	}

	return true, nil
}
