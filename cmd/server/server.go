package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"gophkeeper/internal/logger"
	"gophkeeper/internal/server/apiserver"
	"gophkeeper/internal/server/config"
	"gophkeeper/internal/server/repository"
	"gophkeeper/internal/server/services"
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
	l = l.Level(zerolog.InfoLevel).With().Timestamp().Logger()

	cfg, err := config.New()
	if err != nil {
		l.Fatal().Msgf("config.New: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	store, err := repository.New(repository.Config{
		ConnDSN:         cfg.DSN,
		MaxConn:         10,
		MaxConnLifeTime: 30 * time.Second,
		MaxConnIdleTime: 30 * time.Second,
		Logger:          l,
	}, l)
	if err != nil {
		l.Fatal().Msgf("repository.New: %w", err)
	}

	auth := services.NewAuthService(ctx, store, cfg.Key, l)
	item := services.NewItemService(ctx, store, l)

	s := apiserver.New(cfg, auth, item, l)
	if err = s.Run(ctx); err != nil {
		l.Fatal().Err(err).Msg("server.Run")
	}
}
