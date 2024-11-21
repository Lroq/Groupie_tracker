package groupie

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Filters struct {
	Name           string
	Image          string
	Dates          Date
	Locations      Location
	CreationDate   int
	Members        []string
	FirstAlbumDate string
}

func filterArtists(creationDate int) ([]Filters, error) {
	artists, err := FetchArtists()
	if err != nil {
		return nil, err
	}

	dates, err := GetDates()
	if err != nil {
		log.Fatal(err)
	}

	locations := GetLocations()

	var filters []Filters
	for i, artist := range artists {
		if artist.CreationDate == creationDate || creationDate == 0 {
			filters = append(filters, Filters{
				Name:           artist.Name,
				Image:          artist.Image,
				Dates:          dates[i],
				Locations:      locations[i],
				Members:        artist.Members,
				CreationDate:   artist.CreationDate,
				FirstAlbumDate: artist.FirstAlbumDate,
			})
		}
	}

	return filters, nil
}

func FilterAlbum(firstAlbum string) ([]Filters, error) {
	artists, err := FetchArtists()
	if err != nil {
		return nil, err
	}

	dates, err := GetDates()
	if err != nil {
		log.Fatal(err)
	}

	locations := GetLocations()

	var filters []Filters
	for i, artist := range artists {
		// extrait l'année de la date du premier album
		year := strings.Split(artist.FirstAlbumDate, "-")[2]

		if year == firstAlbum {
			filters = append(filters, Filters{
				Name:           artist.Name,
				Image:          artist.Image,
				Dates:          dates[i],
				Locations:      locations[i],
				Members:        artist.Members,
				CreationDate:   artist.CreationDate,
				FirstAlbumDate: year,
			})
		}
	}

	return filters, nil
}

func FilterLocation(location string) ([]Filters, error) {
	artists, err := FetchArtists()
	if err != nil {
		return nil, err
	}

	dates, err := GetDates()
	if err != nil {
		log.Fatal(err)
	}

	locations := GetLocations()

	var filters []Filters
	added := make(map[string]bool) // garde une trace des artistes déjà ajoutés
	for i, artist := range artists {
		if added[artist.Name] {
			continue
		}

		for _, loc := range locations[i].Locations {
			// regarde si la location contient un tiret
			if strings.Contains(loc, "-") {
				// extrait le pays de la location
				country := strings.Split(loc, "-")[1]

				// convertit les deux chaînes en minuscules pour les comparer
				if strings.EqualFold(country, location) {
					filters = append(filters, Filters{
						Name:           artist.Name,
						Image:          artist.Image,
						Dates:          dates[i],
						Locations:      Location{Locations: []string{loc}}, // Only include the matching location
						Members:        artist.Members,
						CreationDate:   artist.CreationDate,
						FirstAlbumDate: strings.Split(artist.FirstAlbumDate, "-")[2],
					})
					added[artist.Name] = true // ajoute l'artiste à la liste des artistes ajoutés
					break
				}
			}
		}
	}

	return filters, nil
}

func FilterMember(members []string) ([]Filters, error) {
	artists, err := FetchArtists()
	if err != nil {
		return nil, err
	}

	dates, err := GetDates()
	if err != nil {
		log.Fatal(err)
	}

	locations := GetLocations()

	var filters []Filters
	for i, artist := range artists {
		// convertit le nombre de membres en chaîne
		numMembers := strconv.Itoa(len(artist.Members))

		// compare le nombre de membres de l'artiste avec le nombre de membres fourni
		for _, member := range members {
			if member == numMembers {
				// ajoute l'artiste à la liste des filtres
				filters = append(filters, Filters{
					Name:           artist.Name,
					Image:          artist.Image,
					Dates:          dates[i],
					Locations:      locations[i],
					Members:        artist.Members,
					CreationDate:   artist.CreationDate,
					FirstAlbumDate: strings.Split(artist.FirstAlbumDate, "-")[2],
				})
				break
			}
		}
	}

	return filters, nil
}

func GlobalFilter(creationDate string) ([]Filters, error) {
	date, err := strconv.Atoi(creationDate)
	if err != nil && creationDate != "" {
		fmt.Println("Error during conversion")
		return nil, err
	}

	return filterArtists(date)
}
