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
	"math/rand"
	"strconv"
	"time"
)
//"Google" : "https://www.googleapis.com/geolocation/v1/geolocate?key=AIzaSyDhdQvs9XLKd7TVYyYX98WWfB1z4VOddko",
/*
Str holds the returned strings from the JSON decoder functions
 */
var Str struct{
	OWL string
	IPaddr string
	Timezone string
	LatLng string
	IpSearch string
	MapData string
	Pokemon string

}
var StrRand string

/*
Channels for handling the goroutines that initiate the GET function of http on the API url
 */
var ipChan = make(chan []byte)
var ipSeachChan = make(chan []byte)
var timeZoneChan = make(chan []byte)
var owlChan = make(chan []byte)
var latLngChan = make(chan []byte)
var pokeChan = make(chan []byte)

/*
API url map. Searchable string identifiers for functionality in loops
 */

var URLS = map[string]string{
	"IP" : "https://api.ipify.org?format=json",
	"IpSearch" : "http://ip-api.com/json/" + Str.IPaddr,
	"Gtimezone" : "https://maps.googleapis.com/maps/api/timezone/json?location=58.1626388,7.9878993&timestamp=1490978678&key=AIzaSyDhdQvs9XLKd7TVYyYX98WWfB1z4VOddko",
	"OWL": "http://api.openweathermap.org/data/2.5/weather?id=6453405&units=metric&appid=a0a5cd928b34063b9443cfea27292270",
	"Pokemon": "http://pokeapi.co/api/v2/pokemon/42/",
}

/*
main starts the application, handles HTTP requests and initiates the appropriate functions
 */
func main() {


	http.HandleFunc("/", homepage)
	http.HandleFunc("/FormattedJson", searchBox)
	http.HandleFunc("/AltSubmit", formInputHandler)
	http.HandleFunc("/maps", maps)
	http.ListenAndServe(":8001", nil)
}

/*
homepage displays the initial index and layout html
 */
func homepage(w http.ResponseWriter, r *http.Request) {


	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "index.html")

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

/*
searchbox handles text input, if blank it loops through the URLS map
Further development needed for specifying each URL if it fits the input
 */
func searchBox(w http.ResponseWriter, r *http.Request) {
	r1 := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(r1)
	StrRand = strconv.FormatInt(rand.Int63n(500),10) + "/"
	//fmt.Println(StrRand)
	URLS["Pokemon"] = "http://pokeapi.co/api/v2/pokemon/" + StrRand
	//fmt.Println(URLS["Pokemon"])
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	//fmt.Println(r.Form) // print information on server side.
	name := r.Form.Get("input")
	//fmt.Println(name)

	if name == "" {
		//for i := 0; i < len(URLS); i++ {
		//go getJSON(URLS)
		for key := range URLS {
			//ipChan <- key

			if key == "IP" {
				i := URLS[key]
				go getJSON(i)
			} else if key == "IpSearch" {
				i := URLS[key]
				go getJSON(i)

			} else if key == "Gtimezone" {
				i := URLS[key]
				//go getJSON(i)
				go getJSON(i)
			} else if key == "OWL" {
				i := URLS[key]
				go getJSON(i)

				//getJSON(fmt.Sprintf(i, Str.LatLng))
			} else if key == "Pokemon" {
				i := URLS[key]
				go getJSON(i)
			}

		}
		/*
	Get the channel data when it is available and input it to variables for further decoding
	 */
		ip := <- ipChan
		ipSearch := <- ipSeachChan
		latLng := <- latLngChan
		timeZ := <- timeZoneChan
		owl := <- owlChan
		pokemon := <- pokeChan

		Str.IPaddr = decoders.DecodeIP(ip)
		Str.IpSearch = decoders.DecodeIpSearch(ipSearch)
		Str.OWL = decoders.DecodeOWL(owl)
		Str.LatLng = decoders.GetIpLatLng(latLng)
		Str.Timezone = decoders.DecodeTimeZone(timeZ)
		Str.Pokemon = decoders.DecodePokemon(pokemon)

		//fmt.Println(Str)
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
	} else {
		/*
		Loads default template, so that if one of the variables are not called they just load default
		if testers sets the template to correct when called if true.
		 */
		ipTemp := path.Join("templates", "dynamicIndex.tmpl")
		ipsTemp := path.Join("templates", "dynamicIndex.tmpl")
		timeTemp := path.Join("templates", "dynamicIndex.tmpl")
		owlTemp := path.Join("templates", "dynamicIndex.tmpl")
		pokeTemp := path.Join("templates", "dynamicIndex.tmpl")

		/*
		Splits the inputs with semicolon and itterates over the split string
		 */
		splitString := strings.Split(name, ";")


		setMap := make(map[string]bool)
		for _, v := range splitString {
			setMap[v] = true
		}
		for key, value := range setMap {
			if value == true {
				i := URLS[key]
				go getJSON(i)
			}
			if value == true && key == "IP" {
				ip := <- ipChan
				Str.IPaddr = decoders.DecodeIP(ip)
				ipTemp = path.Join("templates", "IP.tmpl")
			}
			if value == true && key == "IpSearch" {
				ipSearch := <- ipSeachChan
				latLng := <- latLngChan
				Str.IpSearch = decoders.DecodeIpSearch(ipSearch)
				Str.LatLng = decoders.GetIpLatLng(latLng)
				ipsTemp = path.Join("templates", "IpSearch.tmpl")
			}
			if value == true && key == "Gtimezone" {
				timeZ := <- timeZoneChan
				Str.Timezone = decoders.DecodeTimeZone(timeZ)
				timeTemp = path.Join("templates", "timezone.tmpl")
			}
			if value == true && key == "OWL" {
				owl := <- owlChan
				Str.OWL = decoders.DecodeOWL(owl)
				owlTemp = path.Join("templates", "owl.tmpl")
			}
			if value == true && key == "Pokemon" {
				pokemon := <- pokeChan
				Str.Pokemon = decoders.DecodePokemon(pokemon)
				pokeTemp = path.Join("templates", "pokemon.tmpl")
			}

			//fmt.Println(Str)

			tp := path.Join("templates", "dynamicIndex.tmpl")

			t, pErr := template.ParseFiles(tp,ipTemp, ipsTemp, timeTemp, owlTemp, pokeTemp)
			if pErr != nil {
				panic(pErr)
			}
			pErr = t.Execute(w, Str)
			if pErr != nil {
				http.Error(w, pErr.Error(), http.StatusInternalServerError)
			}
		}
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

/*
getJSON does a get request to corresponding URL to the URLS map and set the content to a channel
 */
func getJSON(url string) {

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
		fmt.Println(" StatusCode ", response.StatusCode)
		/*
		hdr := response.Header

		for key, value := range hdr {
			fmt.Println(" ", key, ":", value,)
		}
		fmt.Println("response Body:", string(contents))
		fmt.Printf("%q", contents)
		*/
		if url == URLS["Gtimezone"] {
			timeZoneChan <- contents
		}
		if url == URLS["OWL"] {
			owlChan <- contents
		}

		if url == URLS["IP"] {
			ipChan <- contents
		}
		if url == URLS["IpSearch"] {
			ipSeachChan <- contents
			latLngChan <- contents

		}
		if url == URLS["Pokemon"]{
			pokeChan <- contents
		}
	}
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


/*
Concept for handling multiple input buttons
###Only concept, not working. Further development needed###
 */
func formInputHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.

	if r.Form.Get("IP;IpSearch;Pokemon") == "JSON_Raw" {
			fmt.Println("check")

			searchBox(w,r)

	} else if r.Form.Get("Dynamic") == "Dynamic_Only" {
		fmt.Println("ABABBABBABABABBA")

	} else if r.Form.Get("RawFormat") == "ShowCode" {
		fmt.Println("This is code")

	}

}
