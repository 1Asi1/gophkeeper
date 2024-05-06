package repository

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `db:"id"`
	Email    string    `db:"email"`
	Password []byte    `db:"password"`
}

type UserRepository interface {
	CreateUser(context.Context, User) (uuid.UUID, error)
	GetUser(context.Context, string) (User, error)
}

func (s *Store) CreateUser(ctx context.Context, user User) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (s *Store) GetUser(ctx context.Context, email string) (User, error) {

	return User{}, nil
}
