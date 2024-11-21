package main

import (
	"groupie/handlers"
	"log"
	"net/http"
)

func setupHandlers() {
	css := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	js := http.FileServer(http.Dir("script"))
	http.Handle("/script/", http.StripPrefix("/script/", js))
}

func main() {
	setupHandlers()
	handlers.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
