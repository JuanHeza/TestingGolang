package main

import (
	"fmt"
	"math"
	geo "github.com/martinlindhe/google-geolocate"


	)

const (
	AK     = "AIzaSyA6wOgQ3bBxBSlhkvhHCGGZgH5UPSqrcwM"
	lat, lng = 26.9132084, -101.4259082

)

func main() {
	ExampleGeocoder()
}

// ExampleGeocoder demonstrates the different geocoding services
func ExampleGeocoder() {
	client := geo.NewGoogleGeo(AK)
	res, _ := client.Geolocate()
	fmt.Println("Lat:",lat,"Lng:",lng)
	fmt.Println(res)
	fmt.Println("Lat:",lat-res.Lat)
	fmt.Println("Lng:",lng-res.Lng)
	fmt.Println(Distance(lat,lng,res.Lat,res.Lng))
	// Output: &{ 40.7127837 -74.0059413}
}
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
  // must cast radius as float to multiply later
	la1 := lat1 * math.Pi / 180
	la2 := lat2 * math.Pi / 180
	D := (lat2-lat1) * math.Pi / 180
	L := (lon2-lon1) * math.Pi / 180

	r := 6371e3 // Earth radius in METERS
	A:=math.Sin(D/2)*math.Sin(D/2)+math.Cos(la1)*math.Cos(la2)*math.Sin(L/2)*math.Sin(L/2)
	C:=2*math.Atan2(math.Sqrt(A),math.Sqrt(1-A))
	// calculate
//	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return r * C//r * math.Asin(math.Sqrt(h))
}