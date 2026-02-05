package config

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port int
	Env  string
}

func Load() *Config {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	if portStr := os.Getenv("PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			cfg.Port = port
		}
	}

	if env := os.Getenv("ENV"); env != "" {
		cfg.Env = env
	}

	return &cfg
}
