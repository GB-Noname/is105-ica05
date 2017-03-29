package decoders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
)

type Pokemon struct {
	Id     float64
	Name   string
	Height float64
	Weight float64
}

var w Pokemon

func DecodePokemon(test []byte) string {
	//fmt.Printf("q", test)
	// Her brukes det kun et utdrag fra data som var i responsen fra OWL
	// For å bruke strøm fra doGet funksjonen, må hele JSON-strukturen
	// defineres; kun Coordinates og Measurements (main) er definert i
	// dette eksemplet
	var buffer bytes.Buffer

	//v := int64(w.Height)
	//s10 := strconv.FormatInt(v, 1)
	//Height10 := []byte("int (base 10):")
	//Height10 = strconv.AppendInt(Height10, -42, 10)

	//h := w.Height
	//if h, err := strconv.ParseInt(h, 10, 64); err == nil {
	//	fmt.Printf("%T, %v\n", h, h)
	//}

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

		buffer.WriteString("\n PokemonId: " + strconv.FormatFloat(w.Id, 'f', 0, 64))
		buffer.WriteString("\n Pokemon Name: " + w.Name)
		//buffer.WriteString("\n Height: " + w.Height + "Weight:" + w.Weight)
		buffer.WriteString("\n Height:" + strconv.FormatFloat(w.Height, 'f', 0, 64))
		buffer.WriteString("\n Weight:" + strconv.FormatFloat(w.Weight, 'f', 0, 64))
		//fmt.Printf("\n Pokemon: ID: %v, Name: %v, Height: %v, Weight: %v \n",
		//w.Id, w.Name, w.Height, w.Weight)
		fmt.Println(buffer.String())
	}
	return buffer.String()
}
