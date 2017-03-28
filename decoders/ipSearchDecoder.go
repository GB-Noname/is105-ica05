package decoders

import (
	"encoding/json"
	"bytes"
	"io"
	"log"
	"fmt"

)

type IpSearch struct {
	Country string
	RegionName string
	Isp string
	City string
	Reverse string
	Mobile bool
	Proxy bool



}
var w IpSearch
func DecodeIpSearch(test []byte) string {
	//fmt.Printf("q", test)
	// Her brukes det kun et utdrag fra data som var i responsen fra OWL
	// For å bruke strøm fra doGet funksjonen, må hele JSON-strukturen
	// defineres; kun Coordinates og Additional (main) er definert i
	// dette eksemplet
	var buffer bytes.Buffer
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
		buffer.WriteString("Information about your IP: " + w.Reverse)

		//fmt.Printf("\n You are in country: %q \n More specific %q in %q \n Your ISP is: %q\n",
			//w.Country, w.City, w.RegionName, w.Isp)
		buffer.WriteString("\n You are in country: " + w.Country + "\n More specific " + w.City+ " in " + w.RegionName)
		buffer.WriteString("\n Your ISP is: "+w.Isp)
		if w.Mobile == true {
			buffer.WriteString("You are on a mobile network")
		}
		if w.Proxy == true {
			buffer.WriteString("You are suing a proxy server")
		}
		fmt.Println(buffer.String())

	}
	return buffer.String()
}
