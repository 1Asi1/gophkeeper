package apiserver

import (
	"fmt"
	"net"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gophkeeper/internal/server/config"
	"gophkeeper/internal/server/repository"
	"gophkeeper/internal/server/services"
	grpcServer "gophkeeper/internal/server/transports/grpc"
	proto "gophkeeper/rpc/gen"
)

type APIServer struct {
	cfg     config.Config
	log     zerolog.Logger
	storage repository.Store
	authS   services.AuthService
	itemS   services.ItemService
}

func New(cfg config.Config, authS services.AuthService, itemS services.ItemService, storage repository.Store, log zerolog.Logger) APIServer {
	return APIServer{
		cfg:     cfg,
		log:     log,
		storage: storage,
		authS:   authS,
		itemS:   itemS,
	}
}

func (s *APIServer) Run() error {
	tlsCredentials, err := credentials.NewServerTLSFromFile("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return fmt.Errorf("credentials.NewServerTLSFromFile: %w", err)
	}

	server := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(grpcServer.AuthInterceptor(s.authS)),
		grpc.StreamInterceptor(grpcServer.AuthStreamInterceptor(s.authS)),
	)

	proto.RegisterGophkeeperGrpcServer(server, grpcServer.NewGophkeeperGrpcService(s.authS, s.itemS))

	listen, err := net.Listen("tcp", s.cfg.Port)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	go func() {
		if err = server.Serve(listen); err != nil {
			s.log.Fatal().Msgf("failed to start server: %v", err)
		}
	}()

	return nil
}
