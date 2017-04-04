package decoders

import (
	"encoding/json"
	"bytes"
	"io"
	"log"
	"fmt"


)

/* TimezoneResult holds the variables pointing to the JSON response, using the 'json' identifiers
 */
type TimezoneResult struct {
	// DstOffset is the offset for daylight-savings time in seconds.
	DstOffset int `json:"dstOffset"`
	// RawOffset is the offset from UTC for the given location.
	RawOffset int `json:"rawOffset"`
	// TimeZoneID is a string containing the "tz" ID of the time zone.
	TimeZoneID string `json:"timeZoneId"`
	// TimeZoneName is a string containing the long form name of the time zone.
	TimeZoneName string `json:"timeZoneName"`
}


/*
Decodes the TimeZone response and return it as a string
 */
func DecodeTimeZone(test []byte) string{
	var buffer bytes.Buffer
	dec := json.NewDecoder(bytes.NewReader(test))
	fmt.Println(dec)
	for {

		var res TimezoneResult

		if err := dec.Decode(&res); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			fmt.Print("TIMEZONE ERROR!")
		}

		buffer.WriteString("Timezone information: ")

		buffer.WriteString("\n Server is in the timezone: " + res.TimeZoneName + "\n" +
		"With the TimeZoneId: " + res.TimeZoneID)
		}
		fmt.Println(buffer.String())

	return buffer.String()
}


