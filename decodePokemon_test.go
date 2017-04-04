package main

import (
	"testing"
	"./decoders"
	"fmt"
)

type Pokemon struct {
	Id float64
}

func TestDecodePokemon(t *testing.T) {
	//var Id float64
	//Id = Average([]float64{1,2})
	//if Id != 1.5 {
		//t.Error("Expected 1.5, got ", Id)


	//const Id, Name, Height, Weight = float64, "golbat", 16, 550
	//got := DecodePokemon.Index(Id, Name, Height, Weight)
	//if got != Id {
		//t.Errorf("Index(%v) = %v; Id %v")
	go getJSON(URLS["Pokemon"])
	//go getJSON("http://pokeapi.co/api/v2/pokemon/")
	annet := <- pokeChan

		//for i := 0; i < len(slice); i++ {
			//fmt.Printf("%X %+q %b\n", slice[i],)


	confMap := map[string]string{}
	for key, value := range annet {
		confMap[string(key)] = string(value)
	}
	//fmt.Println(confMap)
	//fmt.Println("###############")
	//fmt.Println(confMap["Id"])
	//fmt.Println("###############")
	//fmt.Println(confMap["id"])
	// And then to find values by key:
	//if value, ok := confMap["key1"]; ok {
		// Found
	//}
	//for key, v := range annet {
	//	if v.key == "PokemonId: " {
			// Found!
	//		fmt.Println(v)
	//	}

	/*
	Setting actual to the length of the string returned by DecodePokemon
	Minimum the string is 53 characters long without the variables converted to string
	If variables are nill or 0 they will not be converted and the string will be shorter
	Hence testing for 60 for safe measure is an adequate test for this scenario. *In case nill or 0 get translated
	 */
	actual := len(decoders.DecodePokemon(annet))
	expected := 60
	fmt.Println(actual)
	if actual < expected{
		t.Errorf("Test failed, expected, longer string!")
	}

	//var v float64
	//v = DecodePokemon([]float64{1,2})
	//if v != 1.5 {
	//	t.Error("Expected 1.5, got ", v)
	//}
}