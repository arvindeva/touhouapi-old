package handlers

import (
	"encoding/json"
	"net/http"
	"os"

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

// loadJSONData loads data from a given filepath and decodes it into the target interface.
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
	var touhous []Touhou
	err := loadJSONData("./data/touhou.json", &touhous)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(touhous)
}

// GetTouhouById extracts the path from the URL of the request and gets the last element of the path to retrieve a Touhou by ID.
//
// w http.ResponseWriter, r *http.Request. Returns nothing.
func GetTouhouById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var touhous []Touhou
	err := loadJSONData("./data/touhou.json", &touhous)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, touhou := range touhous {
		if touhou.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(touhou)
			return
		}
	}

	http.NotFound(w, r)
}
