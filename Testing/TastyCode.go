// Example of creating a web based application purely using
// the net/http package to display weather information and
// Twitter Bootstrap so it doesn't look like it's '92.
//
// To start the app, run:
//    go run weatherweb.go
//
// Accessible via:  http://localhost:8888/here
package main

import (
	 "encoding/json"
	// "html/template"

	// owm "github.com/briandowns/openweathermap"
	//	"io/ioutil"
	"os"
	"log"
	// "net/http"
)

// URL is a constant that contains where to find the IP locale info
// Data will hold the result of the query to get the IP
// address of the caller.
type Data struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	ORG         string  `json:"org"`
	AS          string  `json:"as"`
	Message     string  `json:"message"`
	Query       string  `json:"query"`
}

var err error 
func getLocation() (*Data) {
	err = nil
	// response, err := http.Get(http://ip-api.com/json)
	// if err != nil {
		// log.Fatal(err)
	// }
	// defer response.Body.Close()
	code, err:=os.Open("cop.txt")
	r := &Data{}
	// if err = json.NewDecoder(response.Body).Decode(&r); err != nil {
	if err = json.NewDecoder(code).Decode(&r); err != nil {
		return nil
	}
	return r
}
func main(){
	X:=getLocation()
	if err != nil{
		log.Fatal(err)
	} 
	log.Println(*X)
}