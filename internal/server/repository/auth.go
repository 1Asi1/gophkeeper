package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

//go:generate mockgen -source=auth.go -destination=./mock/auth.go -package=mock

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
	id := uuid.New()
	user.ID = id

	query := `
		INSERT INTO tbl_users(id,email,password)
		VALUES (:id, :email, :password)`

	result, err := s.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("CreateUser: %w", err)
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

func (s *Store) GetUser(ctx context.Context, email string) (User, error) {
	query := `
	SELECT
	    id,
		password
	FROM tbl_users
	WHERE email = $1
`
	var model User
	err := s.db.GetContext(ctx, &model, query, email)
	if err != nil {
		return User{}, fmt.Errorf("GetUser: %w", err)
	}

	return model, nil
}
