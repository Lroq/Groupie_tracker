package handlers

import (
	"groupie/src"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/listartist", ListArtist)
	http.HandleFunc("/artist/", groupie.ArtistPage)
}
