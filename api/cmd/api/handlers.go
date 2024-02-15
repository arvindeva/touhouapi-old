package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arvindeva/touhouapi/api/internal/data"
	"github.com/julienschmidt/httprouter"
)

func (app *application) getTouhous(w http.ResponseWriter, r *http.Request) {
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

func (app *application) getTouhouByID(w http.ResponseWriter, r *http.Request) {
	// When httprouter is parsing a request, the values of any named parameters
	// will be stored in the request context. We'll talk about request context
	// in detail later in the book, but for now it's enough to know that you can
	// use the ParamsFromContext() function to retrieve a slice containing these
	// parameter names and values like so:
	params := httprouter.ParamsFromContext(r.Context())

	// We can then use the ByName() method to get the value of the "id" named
	// parameter from the slice and validate it as normal.
	id := params.ByName("id")
	app.logger.Info("HELLO")
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
