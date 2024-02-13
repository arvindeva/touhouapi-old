package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/arvindeva/touhouapi/api/data"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Touhou Project API"))
}

// Event represents the structure of our resource
type Touhou struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Touhous []*Touhou

func loadJSONData(filepath string, target interface{}) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func GetTouhous(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Getting Touhous"))
	touhous, err := data.LoadTouhousJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	touhous.ToJSON(w)
}

func GetTouhouById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	touhous, err := data.LoadTouhousJSON()
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
