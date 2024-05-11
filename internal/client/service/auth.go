package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gophkeeper/internal/client/models"
	proto "gophkeeper/rpc/gen"
)

type AuthService interface {
	Register(context.Context, string, string) error
	Login(context.Context, string, string) error
}

type auth struct {
	log              zerolog.Logger
	client           proto.GophkeeperGrpcClient
	refreshTokenOnce sync.Once
	token            *models.Token
}

func NewAuthService(client proto.GophkeeperGrpcClient, token *models.Token, log zerolog.Logger) AuthService {
	return &auth{
		log:    log,
		client: client,
		token:  token,
	}
}

func (s *auth) Register(ctx context.Context, email string, password string) error {
	tokenData, err := s.client.Register(ctx, &proto.AuthRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok && e.Code() == codes.AlreadyExists {
			return fmt.Errorf("status.FromError: %v", e.Message())
		}
		return fmt.Errorf("s.client.Register: %w", err)
	}

	s.token.Set(tokenData.Token)

	return nil
}

func (s *auth) Login(ctx context.Context, email string, password string) error {
	tokenData, err := s.client.Login(
		ctx,
		&proto.AuthRequest{
			Email:    email,
			Password: password,
		},
	)
	if err != nil {
		if e, ok := status.FromError(err); ok && e.Code() == codes.NotFound {
			return fmt.Errorf("status.FromError: %v", e.Message())
		}

		return fmt.Errorf("s.client.Login: %w", err)
	}

	s.token.Set(tokenData.Token)
	return nil
}
