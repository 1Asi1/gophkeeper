package config

import (
	_ "embed"
	"encoding/json"
	"flag"
	"os"

	"github.com/rs/zerolog"
)

type FileConfig struct {
	Port string `json:"port"`
	Key  string `json:"key"`
	DSN  string `json:"dsn"`
}

type Config struct {
	Port string
	Key  string
	DSN  string
}

func New(l zerolog.Logger) (Config, error) {
	cfgFile := flag.String("c", "config.json", "config name")
	port := flag.String("p", "", "port to run server")
	key := flag.String("k", "", "secret key")
	dsn := flag.String("d", "", "database dsn")

	var cfgPathName string
	cfgFileEnv, ok := os.LookupEnv("CONFIG")
	if ok {
		l.Info().Msgf("config value: %s", cfgFileEnv)
		cfgPathName = cfgFileEnv
	} else {
		l.Info().Msgf("config address value: %s", *cfgFile)
		cfgPathName = *cfgFile
	}

	file, err := os.OpenFile("internal/server/config/"+cfgPathName, os.O_RDONLY, 0644)
	if err != nil {
		return Config{}, err
	}
	defer func() { _ = file.Close() }()

	var cfgFileData FileConfig
	if err = json.NewDecoder(file).Decode(&cfgFileData); err != nil {
		return Config{}, err
	}

	var cfg Config

	portEnv, ok := os.LookupEnv("PORT")
	if ok {
		cfg.Port = portEnv
	} else {
		cfg.Port = *port
		if cfg.Port == "" {
			cfg.Port = cfgFileData.Port
		}
	}

	keyEnv, ok := os.LookupEnv("KEY")
	if ok {
		cfg.Key = keyEnv
	} else {
		cfg.Key = *key
		if cfg.Key == "" {
			cfg.Key = cfgFileData.Key
		}
	}

	dsnEnv, ok := os.LookupEnv("DSN")
	if ok {
		cfg.DSN = dsnEnv
	} else {
		cfg.DSN = *dsn
		if cfg.DSN == "" {
			cfg.DSN = cfgFileData.DSN
		}
	}

	return cfg, nil
}
