package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"gophkeeper/internal/server/models"
	"gophkeeper/internal/server/oops"
)

type AuthService interface {
	Generate(uuid.UUID, time.Time) (string, error)
	GetUser(context.Context) (uuid.UUID, error)
}

type Auth struct {
	Key string
}

func NewAuthService(key string) *Auth {
	return &Auth{key}
}

func (s *Auth) Generate(id uuid.UUID, expireAt time.Time) (string, error) {
	claims := &models.Auth{
		ID:               id,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expireAt)},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.Key))
	if err != nil {
		return "", fmt.Errorf("jwt.NewWithClaims: %w", err)
	}

	return token, nil
}

func (s *Auth) GetUser(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.UUID{}, errors.New("failed to read request metadata")
	}

	var token string
	if values := md.Get("token"); len(values) == 0 {
		return uuid.UUID{}, fmt.Errorf("md.Get: %w", oops.ErrTokenNotFound)
	} else {
		token = values[0]
	}

	data, err := jwt.ParseWithClaims(token, &models.Auth{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Key), nil
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("jwt.ParseWithClaims: %w", err)
	}

	claims, ok := data.Claims.(*models.Auth)
	if !ok && !data.Valid {
		return uuid.UUID{}, nil
	}

	return claims.ID, nil
}
