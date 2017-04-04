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

	go getJSON(URLS["Pokemon"])

	annet := <- pokeChan

	confMap := map[string]string{}
	for key, value := range annet {
		confMap[string(key)] = string(value)
	}

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
}