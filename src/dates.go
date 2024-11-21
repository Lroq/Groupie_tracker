package groupie

import (
	"encoding/json"
	"net/http"
)

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func GetDates() ([]Date, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dates struct {
		Index []Date `json:"index"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&dates); err != nil {
		return nil, err
	}

	return dates.Index, nil
}
