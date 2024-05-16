package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"gophkeeper/internal/server/services"
)

const (
	UserID         = "userID"
	registerMethod = "/gophkeeper.Auth/Register"
	loginMethod    = "/gophkeeper.Auth/Login"
)

var method = map[string]struct{}{
	registerMethod: {},
	loginMethod:    {},
}

type Stream struct {
	grpc.ServerStream
	Ctx context.Context
}

func AuthInterceptor(auth services.AuthService) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		if _, ok := method[info.FullMethod]; !ok {
			userId, err := auth.GetUser(ctx)
			if err != nil {
				return nil, fmt.Errorf("auth.GetUser: %w", err)
			}
			ctxWithUserId := context.WithValue(ctx, UserID, userId)
			return handler(ctxWithUserId, req)
		}
		return handler(ctx, req)
	}
}

func AuthStreamInterceptor(auth services.AuthService) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		if _, ok := method[info.FullMethod]; !ok {
			userId, err := auth.GetUser(ss.Context())
			if err != nil {
				return fmt.Errorf("auth.GetUser: %w", err)
			}
			ctxWithUserId := context.WithValue(ss.Context(), UserID, userId)
			ss.Context()
			return handler(srv, &Stream{
				ServerStream: ss,
				Ctx:          ctxWithUserId,
			})
		}

		return handler(srv, ss)
	}
}
