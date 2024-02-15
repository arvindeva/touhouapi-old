package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/arvindeva/touhouapi/api/internal/data"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type application struct {
	logger  *slog.Logger
	touhous *data.TouhouModel
	router  *mux.Router
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
		logger:  logger,
		touhous: &data.TouhouModel{},
	}

	app.router = app.routes()

	err := http.ListenAndServe(":"+port, app.router)
	app.logger.Error(err.Error())
	os.Exit(1)
}
