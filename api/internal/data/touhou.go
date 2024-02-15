package data

import (
	"github.com/arvindeva/touhouapi/api/utils"
)

type Touhou struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TouhouModel struct{}

func (m *TouhouModel) GetTouhouByID(id string) (Touhou, error) {
	var touhous []*Touhou
	err := utils.LoadJSONData("./internal/data/touhou.json", &touhous)
	if err != nil {
		return Touhou{}, err
	}

	for _, touhou := range touhous {
		if touhou.ID == id {
			return *touhou, nil
		}
	}
	return Touhou{}, ErrNoRecord
}

func (m *TouhouModel) GetTouhous() ([]*Touhou, error) {
	var touhous []*Touhou
	err := utils.LoadJSONData("./internal/data/touhou.json", &touhous)
	if err != nil {
		return nil, err
	}
	return touhous, nil
}
