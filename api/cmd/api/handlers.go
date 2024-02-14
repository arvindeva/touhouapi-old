package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) GetTouhous(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Handle GET Touhous")
	touhous, err := LoadTouhousJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	touhous.ToJSON(w)
}

func (app *application) GetTouhouByID(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Handle GET Touhou by ID")
	vars := mux.Vars(r)
	id := vars["id"]
	touhous, err := LoadTouhousJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, touhou := range touhous {
		if touhou.ID == id {
			w.Header().Set("Content-Type", "application/json")
			touhou.ToJSON(w)
			return
		}
	}

	http.NotFound(w, r)
}
