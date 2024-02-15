package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to the Touhou API Project")
	})
	router.HandleFunc("/touhou", app.GetTouhous)
	router.HandleFunc("/touhou/{id}", app.GetTouhouByID)

	return router
}
