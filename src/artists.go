package groupie

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Image          string   `json:"image"`
	Location       string   `json:"location"`
	Date           string   `json:"date"`
	CreationDate   int      `json:"creationDate"`
	Members        []string `json:"members"`
	FirstAlbumDate string   `json:"firstAlbum"`
}

func FetchArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
