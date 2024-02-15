package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arvindeva/touhouapi/api/internal/data"
	"github.com/gorilla/mux"
)

func (app *application) GetTouhous(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Handle GET Touhous")
	touhous, err := app.touhous.GetTouhous()
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(touhous)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) GetTouhouByID(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Handle GET Touhou by ID")
	vars := mux.Vars(r)
	id := vars["id"]
	touhou, err := app.touhous.GetTouhouByID(id)
	if err != nil {
		app.logger.Error(err.Error())
		if errors.Is(err, data.ErrNoRecord) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(touhou)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
