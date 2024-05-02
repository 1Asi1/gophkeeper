package apiserver

import (
	"github.com/rs/zerolog"
	"gophkeeper/internal/server/config"
)

type APIServer struct {
	cfg config.Config
	log zerolog.Logger
}

func New(cfg config.Config, log zerolog.Logger) APIServer {
	return APIServer{
		cfg: cfg,
		log: log,
	}
}

func (s *APIServer) Run() error {

	return nil
}
