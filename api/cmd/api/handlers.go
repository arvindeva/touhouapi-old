package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/arvindeva/touhouapi/api/internal/data"
	"github.com/julienschmidt/httprouter"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) getTouhous(w http.ResponseWriter, r *http.Request) {
	touhous, err := app.touhou.GetTouhous()
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

	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 1 {
		app.logger.Error("Error converting id to integer:", err)
		app.notFound(w, r)
		return
	}

	touhou, err := app.touhou.GetTouhouByID(id)
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
