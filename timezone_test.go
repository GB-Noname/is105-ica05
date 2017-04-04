package main

import (
	"testing"
	"fmt"
	"./decoders"
)

func TestTimezone(t *testing.T) {
	go getJSON(URLS["Gtimezone"])

	XXX := <- timeZoneChan

	confMap := map[string]string{}
	for key, value := range XXX {
		confMap[string(key)] = string(value)
	}

	actual := len(decoders.DecodeTimeZone(XXX))
	expected := 50
	fmt.Println(actual)
	if actual < expected{
		t.Errorf("Test failed, expected longer string!")
	}
}


