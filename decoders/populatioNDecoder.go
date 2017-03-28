package decoders

import (
	"encoding/json"
	"bytes"
	"io"
	"log"
	"fmt"

)

type Population struct {
	Coord Coordinates
	Main Measurements
}

func DecodePopulation(test []byte) {
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
		var w Weather
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
			w.Coord.Lon, w.Coord.Lat)
		fmt.Printf("Temperature: %.f\n", w.Main.Temp)
		fmt.Printf("Lowest temperature: %.f\n", w.Main.Temp_min)
		fmt.Printf("Peak temperature: %.f\n", w.Main.Temp_max)

	}
}
