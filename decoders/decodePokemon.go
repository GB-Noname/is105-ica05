package decoders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
)

/*
Holds the variables for Pokemon JSON response decoding
 */
type Pokemon struct {
	Id     float64
	Name   string `json:"name"`
	Height float64
	Weight float64
}

var poke Pokemon

/*
Decodes and returns the pokemon information in form of a string
 */
func DecodePokemon(test []byte) string {

	var buffer bytes.Buffer

	dec := json.NewDecoder(bytes.NewReader(test))
	for {

		if err := dec.Decode(&poke); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		buffer.WriteString("\n PokemonId: " + strconv.FormatFloat(poke.Id, 'f', 0, 64))
		buffer.WriteString("\n Pokemon Name: " + poke.Name)
		buffer.WriteString("\n Height:" + strconv.FormatFloat(poke.Height, 'f', 0, 64))
		buffer.WriteString("\n Weight:" + strconv.FormatFloat(poke.Weight, 'f', 0, 64))
		//fmt.Printf("\n Pokemon: ID: %v, Name: %v, Height: %v, Weight: %v \n",
		//w.Id, w.Name, w.Height, w.Weight)
		fmt.Println(buffer.String())
	}
	return buffer.String()
}
