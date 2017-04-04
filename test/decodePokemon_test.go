package test

import (
	"testing"
)



func TestDecodePokemon(t *testing.T) {
	//var Id float64
	//Id = Average([]float64{1,2})
	//if Id != 1.5 {
		//t.Error("Expected 1.5, got ", Id)




	//const Id, Name, Height, Weight = float64, "golbat", 16, 550
	//got := DecodePokemon.Index(Id, Name, Height, Weight)
	//if got != Id {
		//t.Errorf("Index(%v) = %v; Id %v")


	expected := "PokemonId: "
	actual := DecodePokemon("P"[]string{"Height, Weight"})
	if actual != expected{
		t.Errorf("Test failed, expected, actual")
	}

	//var v float64
	//v = DecodePokemon([]float64{1,2})
	//if v != 1.5 {
	//	t.Error("Expected 1.5, got ", v)
	//}
}