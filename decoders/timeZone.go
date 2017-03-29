package decoders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

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
		var w LatLng
		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.PrintF("\n Daylight savings-time in seconds: %.2f\n",
			w.DstOffset)
		fmt.Printf("Offset from UTC in seconds: %.2f", w.RawOffset)
		fmt.Printf("Timezoneid: %q", w.TimeZoneID)
		fmt.Printf("Timezonename: %q", w.TimeZoneName)
	}
}
