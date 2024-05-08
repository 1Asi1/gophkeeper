package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
	"gophkeeper/internal/server/models"
	"gophkeeper/internal/server/oops"
	"gophkeeper/internal/server/repository"
)

type AuthService interface {
	Generate(uuid.UUID, time.Time) (string, error)
	GetUser(context.Context) (uuid.UUID, error)
	Create(context.Context, models.User) (uuid.UUID, error)
	Get(context.Context, string) (models.User, error)
	ComparePassword(models.User, string) (bool, error)
}

type Auth struct {
	log     zerolog.Logger
	storage *repository.Store
	ctx     context.Context
	Key     string
}

func NewAuthService(ctx context.Context, storage *repository.Store, key string, log zerolog.Logger) *Auth {
	return &Auth{
		log:     log,
		storage: storage,
		ctx:     ctx,
		Key:     key,
	}
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

func (s *Auth) Create(ctx context.Context, req models.User) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}

	user := repository.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	id, err := s.storage.CreateUser(ctx, user)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("s.storage.Create: %w", err)
	}

	return id, nil
}

func (s *Auth) Get(ctx context.Context, email string) (models.User, error) {
	res, err := s.storage.GetUser(ctx, email)
	if err != nil {
		return models.User{}, fmt.Errorf("s.storage.Get: %w", err)
	}

	return models.User{
		Email:    res.Email,
		Password: string(res.Password),
	}, err
}

func (s *Auth) ComparePassword(req models.User, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(password)); err != nil {
		return false, fmt.Errorf(" bcrypt.CompareHashAndPassword: %w", err)
	}

	return true, nil
}
