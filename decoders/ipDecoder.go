package decoders

import (
	"encoding/json"
	"bytes"
	"io"
	"log"
	"fmt"

)

/*
Single variable pointer for IP decoding of JSON response
 */
type IP struct {
	Ip string
}

/*
Decodes and returns IP address in form of a string
 */
func DecodeIP(test []byte) string{

	var w IP

	dec := json.NewDecoder(bytes.NewReader(test))
	for {

		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n ip address is: %q\n", w.Ip)

	}
	return w.Ip
}
