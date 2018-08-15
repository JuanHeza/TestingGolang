package main

import (
	"html/template"
	"encoding/json"
	"log"
	"net/http"
    "time"
	"fmt"
	"github.com/skratchdot/open-golang/open"

)

type Device struct {
	ID , Access	string
	Lat, Long	float64
	CO, TEMP, PM10, PM25	string 
}

type ResultN struct {
	Result json.Number `json:"result,Number"`
}

const MainPage string = "https://api.spark.io/v1/devices/"

var List = []Device {
		//Device{"ID","ACCESS","LAT","LONG","CO","TEMP","PM10","PM25",}
		Device{"42003c001551353531343431","2a81f8cf11e752ce750896d2bdb8568db578dc78" ,25.7401114,-100.30569330000003 ,"","","",""},				//1
		//Device{"3d0030001551353531343431","f533f337f0bb3c19f399d306eaf28b8d10df3205" ,25.799225890329403,-100.32385401427746 ,"","","",""},		//2
		// Device{"260037001551353531343431","07ed7ac0d667c3cce4247fb3e786281264db4a22" ,25.733638,-100.301350 ,"","","",""},						//3
		// Device{"3a0021001147343438323536","99efebd2a8b2e81b6a614299a0b94ab1fb0e7705" ,25.803672,-100.332221 ,"","","",""},						//4
		// Device{"2c0042001047343438323536","40cda1e0b8014b3a7824d704ad385baba2b1ebd6" ,25.7616735,-100.31647709999999 ,"","","",""},				//5
		// Device{"38003d001847343438323536","e4bb5ae408a071b39676dfe5974cc8e885d29a63" ,25.613943,-100.264353 ,"","","",""},						//6
		// Device{"570038000b51353335323535","1ce4610d0907aa979a7bc0a20ba370230851de20" ,25.744335,-100.295425 ,"","","",""},						//7
		// Device{"440048000751353530373132","d0843cc0edcce294a318d96b22727260b62500da" ,25.7356122,-100.266857 ,"","","",""},						//8
		// Device{"2b002c001147343438323536","2b3968207c0f70ce7fa77f21557ce4582e52df66" ,25.734130,-100.292794 ,"","","",""},						//9
		Device{"440057000f51353532343635","799b81c43e8167492ecd6a581bbc8d8e8745521f" ,25.7295847,-100.3445827 ,"","","",""},					//10
		// Device{"3a0040000651353530373132","b714d45d14d9010e8728228c7dae2d055709f330" ,25.807721,-100.331524 ,"","","",""},						//11
		Device{"2e0030001447353136383631","971de6af15d978e802478b0d02d505410c751376" ,25.657999,-100.120937 ,"","","",""},						//12
		// Device{"2a0033001247343438323536","83814fe3e3f4a2b41096c134d734a46dc9a60acf" ,25.698499,-100.353065 ,"","","",""},						//13
		Device{"460054000e51353532343635","1d1f08b783105b5d5c383f015c50caf56fc015b8" ,25.7369299,-100.26183549999999 ,"","","",""},				//14
		Device{"220039001947343438323536","db0566f5a8072bac9fa1a0fb3452c3ae93ccc66b" ,25.68386,-100.2467632 ,"","","",""},						//15
		// Device{"2b0024001347343438323536","1fc2844ed5f7868d373a59e760b45db50833ff61" ,25.733475,-100.219216 ,"","","",""},						//16
		Device{"2b001c001347343438323536","4361d138e70958ab716bd8ba3d20697c68a6061f" ,25.667777,-100.5361814 ,"","","",""},					//17
		// Device{"3b0037001147353236343033","cde136ecc36e51d6d8a7376422e3e9619520d2b5" ,25.731426,-100.238799 ,"","","",""},						//18
		// Device{"2f0044000b47343138333038","2ce434c723e206d10c7ba0a380d891d02d376f51" ,25.7893335,-100.0507866 ,"","","",""},					//19
		// Device{"26002f001147343438323536","2ce434c723e206d10c7ba0a380d891d02d376f51" ,25.746714028822527,-100.2875804901123 ,"","","",""},		//20
	}	

func main() {
	log.Println("Start")
	http.HandleFunc("/", Mapping)
	log.Println("Handling")
	log.Fatal(http.ListenAndServe(":12345", nil))
	log.Println("Opening")
	open.StartWith("http://localhost:8080","Chrome")

}

func getData() []Device{
	var Quack []Device
	log.Println("Getting")
	for i,photon:= range List{
		Id:=photon.ID
		Access:=photon.Access
		go CO(&photon,Id,Access)
		go PM10(&photon,Id,Access)	
		go PM25(&photon,Id,Access)	
		go TEMP(&photon,Id,Access)
		time.Sleep(4000 * time.Millisecond)	
		Quack = append(Quack,photon)
		fmt.Println("Photon #",i+1," A:",photon.CO," B:",photon.PM10," C:",photon.PM25," D:",photon.TEMP)
	}
	log.Println("Checked")
	return Quack
}

func Mapping(w http.ResponseWriter, r *http.Request){
	Data := getData()
	t, err := template.ParseFiles("final.html") 
    if err != nil {
  	  log.Print("template parsing error: ", err) 
  	}
    err = t.Execute(w, Data) 
    if err != nil { 
  	  log.Print("template executing error: ", err) 
  	}
	log.Println("Parsed")	
	time.Sleep(12000 * time.Millisecond)	

}

func CO(photon *Device,ID string,access string){
	time.Sleep(100 * time.Millisecond)
	rsp, err := http.Get(MainPage + ID + "/CO?access_token=" + access)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()
	R := &ResultN{}
	if err = json.NewDecoder(rsp.Body).Decode(&R); err != nil {
		log.Println(err)
	}
	photon.CO = string(R.Result)
}

func PM10(photon *Device,ID string,access string){
	time.Sleep(100 * time.Millisecond)
	rsp, err := http.Get(MainPage + ID + "/PM10?access_token=" + access)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()
	R := &ResultN{}
	if err = json.NewDecoder(rsp.Body).Decode(&R); err != nil {
		log.Println(err)
	}
	photon.PM10 = string(R.Result)
}

func PM25(photon *Device,ID string,access string){
	time.Sleep(100 * time.Millisecond)
 	rsp, err := http.Get(MainPage + ID + "/PM25?access_token=" + access)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()
	R := &ResultN{}
	if err = json.NewDecoder(rsp.Body).Decode(&R); err != nil {
		log.Println(err)
	}
	photon.PM25 = string(R.Result)
}

func TEMP(photon *Device,ID string,access string){
	time.Sleep(100 * time.Millisecond)
 	rsp, err := http.Get(MainPage + ID + "/Temp?access_token=" + access)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()
	R := &ResultN{}
	if err = json.NewDecoder(rsp.Body).Decode(&R); err != nil {
		log.Println(err)
	}
	photon.TEMP = string(R.Result)
}