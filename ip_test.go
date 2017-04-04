package main

import (
	"testing"

	"bytes"
)


	type IP struct {
	Ip string
}

func TestIP(t*testing.T) {


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