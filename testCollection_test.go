package main

import (
	"testing"
	"./decoders"
	"fmt"
	"bytes"
)

type Pokemon struct {
	Id float64
}
type IP struct {
	Ip string
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

func TestIP(t*testing.T) {
	/*
	Gets IP request from API and checks if the IP is returned in a single "" string which would mean the request failed
	If string is not equal to "" the test passes
	 */
	go getJSON(URLS["IP"])
	ip := <- ipChan
	//ip :=[]byte{'1','5','8','.','3','7','.','2','4','0','.','6','2'}
	pi := []byte{'"','"'}
	//sz :=len(ip)
	expected:= ip
	actual:= pi
	if bytes.Equal(ip,pi) {
		t.Errorf("Test failed", expected, actual)

	}

}

func TestMultiple(t *testing.T) {
	counter := 0
	for count := 0 ; count <= 25; count++ {
		for key := range URLS {
			go getJSON(URLS[key])
		}
		counter = count
	}

	/*
	Setting actual to the length of the string returned by DecodePokemon
	Minimum the string is 53 characters long without the variables converted to string
	If variables are nill or 0 they will not be converted and the string will be shorter
	Hence testing for 60 for safe measure is an adequate test for this scenario. *In case nill or 0 get translated
	 */
	actual := counter
	expected := 25
	fmt.Println(actual)
	if actual < expected{
		t.Errorf("Test failed, expected, longer string!")
	}
}