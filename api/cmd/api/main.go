package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
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

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to the Touhou API Project")
	})
	router.HandleFunc("/touhou", app.GetTouhous)
	router.HandleFunc("/touhou/{id}", app.GetTouhouByID)

	err := http.ListenAndServe(":"+port, router)
	l.Error(err.Error())
	os.Exit(1)
}
