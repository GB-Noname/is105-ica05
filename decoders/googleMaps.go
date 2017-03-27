package decoders

import (
	"encoding/json"
	"bytes"
	"io"
	"log"
	"fmt"

)

type Location struct {
	Lng float64
	Lat float64
}

type Maps struct {
	Loc Location
}

// GeocodingResult is a single geocoded address
type GeocodingResult struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAddress  string             `json:"formatted_address"`
	Geometry          AddressGeometry    `json:"geometry"`
	Types             []string           `json:"types"`
	PlaceID           string             `json:"place_id"`
}

// AddressComponent is a part of an address
type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// AddressGeometry is the location of a an address
type AddressGeometry struct {
	Location     LatLng       `json:"location"`
	LocationType string       `json:"location_type"`
	Types        []string     `json:"types"`
}
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func GogleDecoder(test []byte) {
	//fmt.Printf("q", test)
	// Her brukes det kun et utdrag fra data som var i responsen fra OWL
	// For å bruke strøm fra doGet funksjonen, må hele JSON-strukturen
	// defineres; kun Coordinates og Measurements (main) er definert i
	// dette eksemplet

	// Definerer en struktur i Golang etter strukturen fra API-en (openweather)
	// Her kan man virkelig se “styrken” av Golangs struct
	// Datafelt i struct må være med en storbokstav og navn må tilsvare
	// de navn som er i jsonStream (de kan begynne med små bokstaver)

	// Ting er strøm-basert, som vi har snakket om tidligere
	dec := json.NewDecoder(bytes.NewReader(test))
	for {
		// Definerer struktur for en instans av Weather strukturen
		// Dette avhenger selvfølgelig om hva som returneres fra
		// webtjenesten (openweather i dette tilfelle)
		var w LatLng
		//var m Measurements
		// Passerer adressen til Weather-strukturen w til funksjonen
		// Decode (som kalles fra en json.NewDecoder med
		// strings.NewReader(jsonStream) som IN-DATA-STRØM
		// Når det ikke er mer data (EOF) bryter vi utførelsen av
		// denne funksjonen med break
		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// Her er et par eksempler på hvordan man kan skrive ut
		// data fra denne webtjenesten på en brukbar måte
		// Dette er noe dere skal prøve å imitere med data
		// fra andre webtjenester (med andre API-er, selvsagt)
		fmt.Printf("\n Coordinates are: longitude %.2f and latitude %.2f\n",
			w.Lat, w.Lng)



	}
}
