package groupie

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type ArtistInfo struct {
	Name           string
	Image          string
	Dates          Date
	Locations      Location
	CreationDate   int
	Members        []string
	FirstAlbumDate string
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	artistName := strings.TrimPrefix(r.URL.Path, "/artist/")
	artistInfo, err := GetArtistInfo(artistName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/artist.html"))
	err = tmpl.Execute(w, artistInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetArtistInfo(name string) (ArtistInfo, error) {
	artists, err := FetchArtists()
	if err != nil {
		return ArtistInfo{}, err
	}

	dates, err := GetDates()
	if err != nil {
		log.Fatal(err)
	}

	locations := GetLocations()

	var artistInfo ArtistInfo
	for i, artist := range artists {
		if artist.Name == name {
			artistInfo = ArtistInfo{
				Name:           artist.Name,
				Image:          artist.Image,
				Dates:          dates[i],
				Locations:      locations[i],
				Members:        artist.Members,
				CreationDate:   artist.CreationDate,
				FirstAlbumDate: artist.FirstAlbumDate,
			}
			break
		}
	}

	return artistInfo, nil
}
