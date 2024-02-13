package data

import (
	"encoding/json"
	"io"
	"log"

	"github.com/arvindeva/touhouapi/api/utils"
)

type Touhou struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Touhous []*Touhou

func (t *Touhou) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	err := e.Encode(t)
	if err != nil {
		log.Println("Error encoding Touhou to JSON:", err)
	}
	return err
}

func (t *Touhous) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	err := e.Encode(t)
	if err != nil {
		log.Println("Error encoding Touhous to JSON:", err)
	}
	return err
}

func LoadTouhousJSON() (Touhous, error) {
	var touhous Touhous
	err := utils.LoadJSONData("./data/touhou.json", &touhous)
	if err != nil {
		return nil, err
	}
	return touhous, nil
}
