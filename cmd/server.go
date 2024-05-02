package cmd

import (
	"fmt"
	"log"

	"github.com/rs/zerolog"
	"gophkeeper/internal/logger"
	"gophkeeper/internal/server/apiserver"
	"gophkeeper/internal/server/config"
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

	cfg, err := config.New(logger.NewLogger())
	if err != nil {
		log.Fatal("config.New")
	}

	l := logger.NewLogger()
	l = l.Level(zerolog.InfoLevel).With().Timestamp().Logger()

	s := apiserver.New(cfg, l)
	if err = s.Run(); err != nil {
		l.Fatal().Err(err).Msg("server.Run")
	}
}
