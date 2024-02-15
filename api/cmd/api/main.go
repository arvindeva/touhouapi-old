package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/arvindeva/touhouapi/api/internal/data"
	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
	touhou *data.TouhouModel
	router http.Handler
}

func init() {
	// Load values from .env file into the system
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found")
	}
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
		touhou: &data.TouhouModel{},
	}

	app.router = app.routes()
	srv := &http.Server{
		Addr:        ":" + port,
		Handler:     app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
	}

	logger.Info("starting server", "addr", srv.Addr)

	err := srv.ListenAndServe()
	app.logger.Error(err.Error())
	os.Exit(1)
}
