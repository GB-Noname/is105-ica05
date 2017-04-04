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
IpSearch holds the variables pointing to the JSON response
 */
type IpSearch struct {
	Country string
	RegionName string
	Isp string
	City string
	Reverse string
	Mobile bool
	Proxy bool
	Lat float64
	Lon float64



}
var w IpSearch
/*
DecodeIpSearch decodes the JSON and return it in the form of a string
 */
func DecodeIpSearch(test []byte) string {

	var buffer bytes.Buffer



	dec := json.NewDecoder(bytes.NewReader(test))
	for {
		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		buffer.WriteString("Information about server IP: " + w.Reverse)

		buffer.WriteString("\n Registered in country: " + w.Country + "\n More specific " + w.City+ " in " + w.RegionName)
		buffer.WriteString("\n The ISP is: "+w.Isp)
		if w.Mobile == true {
			buffer.WriteString("You are on a mobile network")
		}
		if w.Proxy == true {
			buffer.WriteString("You are suing a proxy server")
		}
		fmt.Println(buffer.String())

	}
	return buffer.String()
}
/*
GetIpLatLng only decodes the latitude and longitude to return it in a single string
 */
func GetIpLatLng(test []byte) string {

	var buffer2 bytes.Buffer



	dec := json.NewDecoder(bytes.NewReader(test))
	for {
		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		buffer2.WriteString(strconv.FormatFloat(w.Lat,'f',7,64))
		buffer2.WriteString("," + strconv.FormatFloat(w.Lon,'f',7,64))
		//buffer.WriteString("&timestamp=1458000000&key=AIzaSyDhdQvs9XLKd7TVYyYX98WWfB1z4VOddko")

		fmt.Println(buffer2.String())

	}
	return buffer2.String()
}

