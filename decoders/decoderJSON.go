package decoders

import (
	"encoding/json"
	"bytes"
	"io"
	"log"
	"fmt"

	"strconv"
)

type Coordinates struct {
	Lon float64
	Lat float64
}
type Measurements struct {
	Temp float64
	Pressure float64
	Humidity float64
	Temp_min float64
	Temp_max float64
}
type Weather struct {
	Coord Coordinates
	Main Measurements
}

func DecodeOWL(test []byte) string{
	//fmt.Printf("q", test)
	// Her brukes det kun et utdrag fra data som var i responsen fra OWL
	// For å bruke strøm fra doGet funksjonen, må hele JSON-strukturen
	// defineres; kun Coordinates og Additional (main) er definert i
	// dette eksemplet

	// Definerer en struktur i Golang etter strukturen fra API-en (openweather)
	// Her kan man virkelig se “styrken” av Golangs struct
	// Datafelt i struct må være med en storbokstav og navn må tilsvare
	// de navn som er i jsonStream (de kan begynne med små bokstaver)
	var buffer bytes.Buffer
	// Ting er strøm-basert, som vi har snakket om tidligere
	dec := json.NewDecoder(bytes.NewReader(test))
	for {
		// Definerer struktur for en instans av Weather strukturen
		// Dette avhenger selvfølgelig om hva som returneres fra
		// webtjenesten (openweather i dette tilfelle)
		var w Weather
		//var m Additional
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

		buffer.WriteString("\nCoordinates are: longitude " + strconv.FormatFloat(w.Coord.Lon,'f',2,64) )
		buffer.WriteString(" and latitude " + strconv.FormatFloat(w.Coord.Lat,'f',2,64) )
		buffer.WriteString("\nTemperature: \n" + strconv.FormatFloat(w.Main.Temp,'f',2,64) )
		buffer.WriteString("\nLowest temperature: \n" + strconv.FormatFloat(w.Main.Temp_min,'f',2,64) )
		buffer.WriteString("\nPeak temperature: \n" + strconv.FormatFloat(w.Main.Temp_max,'f',2,64) )



		fmt.Println(buffer.String())

	}
	return buffer.String()
}
