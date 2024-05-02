package config

import "github.com/rs/zerolog"

type Config struct {
}

func New(l zerolog.Logger) (Config, error) {
	return Config{}, nil
}
