package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/miank1/GreenlightAPI/internal/config"
)

type application struct {
	config config.Config
	logger *log.Logger
}

func main() {
	fmt.Println("starting Greenlight API")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error { // Read this once
	// Load config
	cfg := config.Load()

	// Initialize logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Create application struct
	app := &application{ // why ??
		config: *cfg,
		logger: logger,
	}

	// Setup HTTP server   // timeout
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)

	return srv.ListenAndServe()
}
