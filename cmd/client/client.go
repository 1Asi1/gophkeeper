package main

import (
	"context"
	"fmt"

	"gophkeeper/internal/client/cli"
	"gophkeeper/internal/client/config"
	"gophkeeper/internal/client/models"
	"gophkeeper/internal/client/service"
	"gophkeeper/internal/client/transports/grpc"
	"gophkeeper/internal/logger"
	proto "gophkeeper/rpc/gen"
)

var (
	BuildVersion = "N/A"
	BuildDate    = "N/A"
	BuildCommit  = "N/A"
)

func main() {
	fmt.Printf("Build version: %s\n", BuildVersion)
	fmt.Printf("Build date: %s\n", BuildDate)
	fmt.Printf("Build commit: %s\n", BuildCommit)

	l := logger.NewLogger()

	cfg, err := config.New()
	if err != nil {
		l.Fatal().Msgf("config.New: %v", err)
	}

	var token *models.Token
	client, err := grpc.CreateGrpcConnection(cfg.ServerADDR, token)
	if err != nil {
		l.Fatal().Msgf("grpc.CreateGrpcConnection error: %v", err)
	}

	auth := service.NewAuthService(proto.NewGophkeeperGrpcClient(client), token, l)
	item := service.NewItemService(proto.NewGophkeeperGrpcClient(client), cfg.Key, l)

	term := cli.New(auth, item, l)
	if err := term.Execute(context.Background()); err != nil {
		l.Error().Msgf("term.Execute: %v", err)
	}
}
