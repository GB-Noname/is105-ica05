package decoders

import (
	"encoding/json"
	"bytes"
	"io"
	"log"
	"fmt"

	"strconv"
)
/*
Coordinates holds the variables pointing to location values in the JSON response
 */
type Coordinates struct {
	Lon float64
	Lat float64
}
/*
Measurements holds the variables pointing to the temperature values in the JSON response
 */
type Measurements struct {
	Temp float64
	Pressure float64
	Humidity float64
	Temp_min float64
	Temp_max float64
}
/*
Weather binds the two structs together
 */
type Weather struct {
	Coord Coordinates
	Main Measurements
}

/*
Decodes the OWL JSON response and returns it in a string using the buffer
 */
func DecodeOWL(test []byte) string{

	var buffer bytes.Buffer
	dec := json.NewDecoder(bytes.NewReader(test))
	for {

		var w Weather

		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		buffer.WriteString("Live weather information about Kristiansand")
		buffer.WriteString("\nCoordinates are: latitude " + strconv.FormatFloat(w.Coord.Lat,'f',7,64) )
		buffer.WriteString(" and longitude " + strconv.FormatFloat(w.Coord.Lon,'f',7,64) )
		buffer.WriteString("\nTemperature: \n" + strconv.FormatFloat(w.Main.Temp,'f',2,64) )
		buffer.WriteString("\nLowest temperature: \n" + strconv.FormatFloat(w.Main.Temp_min,'f',2,64) )
		buffer.WriteString("\nPeak temperature: \n" + strconv.FormatFloat(w.Main.Temp_max,'f',2,64) )

		fmt.Println(buffer.String())

	}
	return buffer.String()
}
