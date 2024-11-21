package handlers

import (
	"groupie/src"
	"html/template"
	"net/http"
)

type PageData struct {
	Artists         []groupie.ArtistInfo
	FilteredArtists []groupie.Filters
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func ListArtist(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	creationDate := r.FormValue("creation_date")
	firstAlbum := r.FormValue("album_date")
	location := r.FormValue("location")
	members := r.Form["members"]

	var data PageData
	var filters []groupie.Filters
	var err error

	if creationDate != "" && creationDate != "1959" {
		filters, err = groupie.GlobalFilter(creationDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if firstAlbum != "" && firstAlbum != "1964" {
		filters, err = groupie.FilterAlbum(firstAlbum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if location != "" && location != "all" {
		filters, err = groupie.FilterLocation(location)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if len(members) > 0 {
		filters, err = groupie.FilterMember(members)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if (creationDate == "" || creationDate == "1959") && (firstAlbum == "" || firstAlbum == "1964") && (location == "" || location == "all") && len(members) == 0 {
		artists, err := groupie.FetchArtists()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var artistInfos []groupie.ArtistInfo
		for _, artist := range artists {
			artistInfo, err := groupie.GetArtistInfo(artist.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			artistInfos = append(artistInfos, artistInfo)
		}
		data = PageData{Artists: artistInfos}
	} else {
		data = PageData{FilteredArtists: filters}
	}

	tmpl.Execute(w, data)
}
