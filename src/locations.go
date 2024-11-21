package groupie

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
)

type Location struct {
	ID        int        `json:"id"`
	Locations []string   `json:"locations"`
	Index     []Location `json:"index"`
}

func GetLocations() []Location {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		log.Fatal(err)
	}
	sort.SliceStable(location.Index, func(i, j int) bool {
		return location.Index[i].ID < location.Index[j].ID
	})
	return location.Index
}
