package decoders

import (
	"encoding/json"
	"bytes"
	"io"
	"log"
	"fmt"


)

// TimezoneResult is a single timezone result.
type TimezoneResult struct {
	// DstOffset is the offset for daylight-savings time in seconds.
	DstOffset int `json:"dstOffset"`
	// RawOffset is the offset from UTC for the given location.
	RawOffset int `json:"rawOffset"`
	// TimeZoneID is a string containing the "tz" ID of the time zone.
	TimeZoneID string `json:"timeZoneId"`
	// TimeZoneName is a string containing the long form name of the time zone.
	TimeZoneName string `json:"timeZoneName"`
}

var buffer bytes.Buffer
func DecodeTimeZone(test []byte) string{
	//fmt.Printf("q", test)
	// Her brukes det kun et utdrag fra data som var i responsen fra OWL
	// For å bruke strøm fra doGet funksjonen, må hele JSON-strukturen
	// defineres; kun Coordinates og Additional (main) er definert i
	// dette eksemplet

	// Definerer en struktur i Golang etter strukturen fra API-en (openweather)
	// Her kan man virkelig se “styrken” av Golangs struct
	// Datafelt i struct må være med en storbokstav og navn må tilsvare
	// de navn som er i jsonStream (de kan begynne med små bokstaver)

	// Ting er strøm-basert, som vi har snakket om tidligere
	dec := json.NewDecoder(bytes.NewReader(test))
	fmt.Println(dec)
	for {
		// Definerer struktur for en instans av Weather strukturen
		// Dette avhenger selvfølgelig om hva som returneres fra
		// webtjenesten (openweather i dette tilfelle)
		var res TimezoneResult

		//var m Additional
		// Passerer adressen til Weather-strukturen w til funksjonen
		// Decode (som kalles fra en json.NewDecoder med
		// strings.NewReader(jsonStream) som IN-DATA-STRØM
		// Når det ikke er mer data (EOF) bryter vi utførelsen av
		// denne funksjonen med break
		if err := dec.Decode(&res); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			fmt.Print("TWTTWTWTW")
		}


		buffer.WriteString("Timezone information: ")

		//fmt.Printf("\n You are in country: %q \n More specific %q in %q \n Your ISP is: %q\n",
		//w.Country, w.City, w.RegionName, w.Isp)
		buffer.WriteString("\n Server is in the timezone: " + res.TimeZoneName + "\n" +
		"With the TimeZoneId: " + res.TimeZoneID)
		}
		fmt.Println(buffer.String())
		// Her er et par eksempler på hvordan man kan skrive ut
		// data fra denne webtjenesten på en brukbar måte
		// Dette er noe dere skal prøve å imitere med data
		// fra andre webtjenester (med andre API-er, selvsagt)
	return buffer.String()
	}


