package data

import (
	"encoding/json"
	"io"

	"github.com/arvindeva/touhouapi/api/utils"
)

type Touhou struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Touhous []*Touhou

func (t *Touhou) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(t)
	return err
}

func (t *Touhous) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(t)
	return err
}

func LoadTouhousJSON() (Touhous, error) {
	var touhous Touhous
	err := utils.LoadJSONData("./internal/data/touhou.json", &touhous)
	if err != nil {
		return nil, err
	}
	return touhous, nil
}
