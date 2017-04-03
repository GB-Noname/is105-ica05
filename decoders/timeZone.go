package decoders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

//
type TimezoneRequest struct {
	Lng float64
	Lat float64
}

type TimezoneResult struct {
	DstOffset    int    `json:"dstOffset"`
	RawOffset    int    `json:"RawOffset"`
	TimeZoneID   string `json:"TimeZoneId"`
	TimeZoneName string `json:"TimeZoneName"`
}

func DecodeTimezone(test []byte) {

	dec := json.NewDecoder(bytes.NewReader(test))
	for {
		var w TimezoneResult
		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Offset for daylight-savings in seconds:")
		fmt.Println(w.DstOffset)
		fmt.Println("Offset for UTC in seconds:")
		fmt.Println(w.RawOffset)
		fmt.Printf("Timezoneid: %v", w.TimeZoneID)
		fmt.Printf("Timezonename: %v", w.TimeZoneName)
	}
}
