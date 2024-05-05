package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Auth struct {
	ID uuid.UUID
	jwt.RegisteredClaims
}
