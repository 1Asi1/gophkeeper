package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

type FileConfig struct {
	ServerADDR string `json:"addr"`
	KeyPath    string `json:"key_path"`
}

type Config struct {
	ServerADDR string
	KeyPath    string
	Key        *rsa.PrivateKey
}

func New() (Config, error) {
	cfgFile := flag.String("c", "config.json", "config name")
	addr := flag.String("a", "", "addresses to server")
	keyPath := flag.String("k", "", "secret key")

	var cfgPathName string
	cfgFileEnv, ok := os.LookupEnv("CONFIG")
	if ok {
		cfgPathName = cfgFileEnv
	} else {
		cfgPathName = *cfgFile
	}

	file, err := os.OpenFile("internal/client/config/"+cfgPathName, os.O_RDONLY, 0644)
	if err != nil {
		return Config{}, err
	}
	defer func() { _ = file.Close() }()

	var cfgFileData FileConfig
	if err = json.NewDecoder(file).Decode(&cfgFileData); err != nil {
		return Config{}, err
	}

	var cfg Config

	addrEnv, ok := os.LookupEnv("ADDR")
	if ok {
		cfg.ServerADDR = addrEnv
	} else {
		cfg.ServerADDR = *addr
		if cfg.ServerADDR == "" {
			cfg.ServerADDR = cfgFileData.ServerADDR
		}
	}

	keyPathEnv, ok := os.LookupEnv("KEY_PATH")
	if ok {
		cfg.KeyPath = keyPathEnv
	} else {
		cfg.KeyPath = *keyPath
		if cfg.ServerADDR == "" {
			cfg.KeyPath = cfgFileData.KeyPath
		}
	}

	pemBytes, err := os.ReadFile(cfg.KeyPath)
	if err != nil {
		return Config{}, fmt.Errorf("os.ReadFile: %w", err)
	}
	block, _ := pem.Decode(pemBytes)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return Config{}, fmt.Errorf("x509.ParsePKCS1PrivateKey: %w", err)
	}
	cfg.Key = key

	return cfg, nil
}
