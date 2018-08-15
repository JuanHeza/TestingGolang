package main//Ororo

import (
	"os"
    "log"
    "fmt"
	"net"
	"time"
	"bytes"
	"strconv"
    "io/ioutil"
    "net/http"
    "github.com/oschwald/geoip2-golang"
    owm "github.com/briandowns/openweathermap"
)

var ip, broadcast net.IP
var Host string
var TimeZone string
var Sunrise, Sunset time.Time
var Units string

func main(){
	Units="C"
	local:=false
	ip,Host,err:=Get_IP(local)
	if err !=nil{
        log.Fatal(err)
	}
	fmt.Print("\tConeccion:\t ")
	log.Println("\n\tHostname:\t",Host,"\n\tIP:\t\t",ip)
	switch (local){
		case true:
			mask := ip.DefaultMask()
			broadcast := net.IP(make([]byte, 4))
			for i := range ip {
				broadcast[i] = ip[i] | (^mask[i])
			}
			fmt.Println("\tBroadcast:\t",broadcast)
		case false:
			X,Y,Z:=LatLong(ip)
			fmt.Println("\tLatitud:\t",X,"\n\tLongitud:\t",Y)
			fmt.Println(Z,"\n")
			ClimaActual(X,Y)
	}
}

func ClimaActual(lat float64, long float64){
	//"github.com/briandowns/openweathermap"
    w, err := owm.NewCurrent(Units, "SP")
    if err != nil {
        log.Fatalln(err)
    }
    w.CurrentByCoordinates(lat,long)
// Rise & Dawn
		loc,_:= time.LoadLocation(TimeZone)
		i, err := strconv.ParseInt(strconv.Itoa(w.Sys.Sunrise), 10, 64)
		if err != nil {
			panic(err)
		}
		Sunrise := time.Unix(i, 0).In(loc)
		fmt.Println("\tSunrise:",Sunrise)  
		
		i, err = strconv.ParseInt(strconv.Itoa(w.Sys.Sunset), 10, 64)
		if err != nil {
			panic(err)
		}
		Sunset := time.Unix(i, 0).In(loc)
		fmt.Println("\tSunset:\t",Sunset)
// Actual
		for i:=range w.Weather{
			weather:=w.Weather[i]
			fmt.Println("\tEstado Actual:\t",weather.Description)
		}  
// Temp & Humidity
		Temp:=fmt.Sprintf("%.1f %s°",w.Main.Temp,Units)
		Humidity:=fmt.Sprint(w.Main.Humidity,"%")
		fmt.Println("\tTemperature:",Temp,"Humidity:",Humidity)
// Wind & Direction
		Wind:=w.Wind.Speed
		Direction:=w.Wind.Deg
		var directions = make(map[string][]float64)
		var direc bool
		var Widi string
		// type Spec struct{
			// A	float64
			// B	float64
			// Ico	rune
		// }
		//Cardinal Degree(Min, Max)
		directions=map[string][]float64{
			"N":	{337.5, 22.5},
			"NE":	 {22.5, 67.5},
			"E":	 {67.5,112.5},
			"SE":	{112.5,157.5},
			"S":	{157.5,202.5},
			"SW":	{202.5,247.5},
			"W":	{247.5,292.5},
			"NW":	{292.5,337.5},
		}
		for key,D:=range directions{
		  A:=D[0]
		  B:=D[1]
			if key=="N"{
				direc = (A<Direction&&Direction<360 || 0<Direction&&Direction<B)
				if direc{
					Widi=fmt.Sprint("\tWind:",Wind,"KM/H\tDirection:",key)
				}
			}else{
				direc = (A<Direction&&Direction<B)
				if direc{
					Widi=fmt.Sprint("\tWind: ",Wind," KM/H\tDirection: ",key)
				}
			}
		}
		fmt.Println(Widi)
}

func LatLong(ip net.IP) (float64,float64,string) {
    //"github.com/oschwald/geoip2-golang"
	db, err := geoip2.Open("GeoLite2-City.mmdb")
    if err != nil {
			log.Println("IP NO ENCONTRADA")
            log.Fatal(err)
    }
    defer db.Close()
    record, err := db.City(ip)
    if err != nil {
            log.Fatal(err)
    }
	Data:=fmt.Sprint("\tDesde\t\t ",record.City.Names["en"],", ",record.Subdivisions[0].Names["en"],", ",record.Country.Names["en"])
/*  
    fmt.Printf("ISO country code: %v\n", record.Country.IsoCode) //	"MX"
*/
    TimeZone = record.Location.TimeZone		 // "American/Monterrey" 

	return record.Location.Latitude, record.Location.Longitude,Data
}

func Get_IP(local bool) (net.IP,string,error){
var ip net.IP
	name, err := os.Hostname()
	if err != nil {
		 fmt.Printf("Oops: %v\n", err)
		 return net.ParseIP("<Nil>"),"<Nil>",err
	}
	switch (local){
		case true:
			addrs, _ := net.LookupIP(name)
			ip=addrs[5].To4()
		case false:
			rsp, err := http.Get("http://checkip.amazonaws.com")
			if err != nil {
				return net.ParseIP("<Nil>"),name, err
				log.Println("IMPOSIBLE CONECTAR")
			}
			defer rsp.Body.Close()

			buf, err := ioutil.ReadAll(rsp.Body)
			if err != nil {
				return net.ParseIP("<Nil>"),name, err
				log.Println("IMPOSIBLE CONECTAR")
			}
			ip=net.ParseIP(string(bytes.TrimSpace(buf))) 
	}

    return ip,name,nil
}