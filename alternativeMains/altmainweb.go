package alternativeMains
import (

	"time"
	"strconv"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"github.com/go-martini/martini"
)

type SiteData struct {
	Name string
	Description string
	Count string
	ServerIP string
	MapData string
}

var count int = 0

func main() {

	m := martini.Classic()

	m.Use( render.Renderer(render.Options{
		IndentJSON: true, // so we can read it..
	}))

	m.Get("/", func(r render.Render, x *http.Request) {
		place := string(x.FormValue("place"))
		place = strings.Replace(place, " ", "+", -1)

		if len(place) <= 0 {place = "UIA+Kristiansand"}
		r.HTML(200, "index", SiteData{"Christian", "Veldig kjekk gutt", strconv.Itoa(count), getServerIP(), place})
	})




	m.Get("/getCount", func() string {
		print(count)
		return strconv.Itoa(count)
	})


	go countNumber()
	m.RunOnAddr(":5050")
	m.Run()


}

func getServerIP() string{
	readApi, err := http.Get("https://api.ipify.org")
	if err != nil {log.Fatal(err)}
	bytes, err := ioutil.ReadAll(readApi.Body)
	if err != nil {log.Fatal(err)}
	return string(bytes)
}

func countNumber()  {
	for true{
		count += 1
		time.Sleep(1 * time.Second)
	}


}