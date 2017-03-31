package main

import (
	"html/template"
	"net/http"
	"path"

	"fmt"
	"./decoders"
	"log"
	"io/ioutil"

	"bytes"
	"strings"
)
//"Google" : "https://www.googleapis.com/geolocation/v1/geolocate?key=AIzaSyDhdQvs9XLKd7TVYyYX98WWfB1z4VOddko",
var Str struct{
	OWL string
	IPaddr string
	Timezone string
	LatLng string
	IpSearch string
	MapData string
}
var JStruct struct {
	Jowl []byte
	jIpaddr []byte
	jTimeZone []byte
	jLatLng []byte
	jIpSearch []byte
}
var ipChan = make(chan []byte )
var ipSeachChan = make(chan []byte)
var timeZoneChan = make(chan []byte)
var owlChan = make(chan []byte)
var latLngChan = make(chan []byte)

//var URLS = make([]string, 3)
var URLS = map[string]string{
	"IP" : "https://api.ipify.org?format=json",
	"IpSearch" : "http://ip-api.com/json/" + Str.IPaddr,
	"Gtimezone" : "https://maps.googleapis.com/maps/api/timezone/json?location=58.1626388,7.9878993&timestamp=1490978678&key=AIzaSyDhdQvs9XLKd7TVYyYX98WWfB1z4VOddko",
	"OWL" : "http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1",
}
var JPS []byte
func main() {

	//fmt.Println(IPaddr)
	//http.HandleFunc("/search", search)
	//URLS[0] = "http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1"
	//URLS[1] = "http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1"
	//URLS[2] = "https://www.googleapis.com/geolocation/v1/geolocate?key=AIzaSyDhdQvs9XLKd7TVYyYX98WWfB1z4VOddko"
	http.HandleFunc("/", homepage)
	http.HandleFunc("/search", searchBox)
	http.HandleFunc("/AltSubmit", formInputHandler)
	http.HandleFunc("/maps", maps)
	//http.HandleFunc("/ttt", searchBox)
	//http.HandleFunc("/ddd", searchBox)
	//http.HandleFunc("/fff", searchBox)
	http.ListenAndServe(":8001", nil)
	//go http.HandleFunc("/", searchBox)
}

func homepage(w http.ResponseWriter, r *http.Request) {


	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "indexTest.html")

	// Note that the layout file must be the first parameter in ParseFiles
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, "test"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func searchBox(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	name := r.Form.Get("name")
	fmt.Println(name)

	if name == "all" {
		//for i := 0; i < len(URLS); i++ {
		//go doGet(URLS)
		for key := range URLS {
			//ipChan <- key
			if key == "Gtimezone" {
			i := URLS[key]
			//go doGet(i)
			go doGet(i)

			} else if key == "OWL" {
				i := URLS[key]
				go doGet(i)
			}else if key == "IP" {
				i := URLS[key]
				go doGet(i)

			} else if key == "IpSearch" {
				i := URLS[key]
				go doGet(i)

			}

				//doGet(fmt.Sprintf(i, Str.LatLng))
			}

	}

	ip := <- ipChan
	ipSearch := <- ipSeachChan
	timeZ := <- timeZoneChan
	owl := <- owlChan
	latLng := <- latLngChan
	Str.Timezone = decoders.DecodeTimeZone(timeZ)
	Str.OWL = decoders.DecodeOWL(owl)
	Str.LatLng = decoders.GetIpLatLng(latLng)
	Str.IPaddr = decoders.DecodeIP(ip)
	Str.IpSearch = decoders.DecodeIpSearch(ipSearch)

	fmt.Println(Str)
	lp := path.Join("templates", "index.tmpl")
	tp := path.Join("templates", "layout.html")
	t, pErr := template.ParseFiles(lp, tp)
	if pErr != nil {
		panic(pErr)
	}
	pErr = t.Execute(w, Str)
	if pErr != nil {
		http.Error(w, pErr.Error(), http.StatusInternalServerError)

	}

}

func maps(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	Str.MapData = r.Form.Get("place")
	newplace := strings.Replace(Str.MapData, " ", "+", -1)
	if len(newplace) <= 0 {Str.MapData = "UIA+Kristiansand"}

	lp := path.Join("templates", "index.tmpl")
	tp := path.Join("templates", "layout.html")
	t, pErr := template.ParseFiles(lp, tp)
	if pErr != nil {
		panic(pErr)
	}
	pErr = t.Execute(w, Str)
	if pErr != nil {
		http.Error(w, pErr.Error(), http.StatusInternalServerError)

	}
}

func doGet(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		fmt.Println(" ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println(" ", key, ":", value,)

		}
		fmt.Println("response Body:", string(contents))
		fmt.Printf("%q", contents)
		if url == URLS["Gtimezone"] {
			//Str.Timezone = decoders.DecodeTimeZone(contents)
			//JStruct.jTimeZone = contents
			timeZoneChan <- contents
		}
		if url == URLS["OWL"] {
			//JStruct.Jowl = contents
			//return contents
			owlChan <- contents
		}

		if url == URLS["IP"] {
			//Str.IPaddr = decoders.DecodeIP(contents)
			//return contents
			//JStruct.jIpaddr = contents

			ipChan <- contents
		}
		if url == URLS["IpSearch"] {


			//go decoders.DecodeIpSearch(contents)
			//JStruct.jIpSearch =contents
			ipSeachChan <- contents
			latLngChan <- contents

		}

		//response.Header.Set("Content-Type", "application/json")
		//go DecodeOWL(js)
	}
	fmt.Printf("Result %s", JStruct)
}

func getGoogle(url string) {
	//response, err := http.Get(url)

	// handle err
	var jsonStr = []byte(`{
  "macAddress": "00:25:9c:cf:1c:ac",
  "signalStrength": -43,
  "age": 0,
  "channel": 11,
  "signalToNoiseRatio": 0
}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	fmt.Printf("%q", body)


	go decoders.GogleDecoder(body)
}



func formInputHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	for k := range r.Form {
		if k == "List" {
		fmt.Println("testetettetetetetet")
			fmt.Println(k)

	} else if k == "ShowProg" {
			fmt.Println("ABABBABBABABABBA")
			fmt.Println(k)
	} else if k == "ShowCode" {
			fmt.Println("This is code")
			fmt.Println(k)
		}
	}
}
