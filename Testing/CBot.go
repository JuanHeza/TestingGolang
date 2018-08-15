package CBot

import (
	"fmt"
	"log"
    "time"
	"bytes"
    "strings"
    "net/http"
    "io/ioutil"
)
const Banned string = "Access Denied. This room has been banned."
const Offline string = "La sala está actualmente fuera de linea"
const MainPage string = "https://es.chaturbate.com/"
const FullVid string = "https://es.chaturbate.com/fullvideo/?b="

// MAKE AN INTERFACE STRING()

type Data struct{
	Status bool
	Changed bool
	In time.Time
	Out time.Time
}

var Actual time.Time
var List = make(map[string]*Data)
var buf bytes.Buffer

//She Has A Big Green Eyes

func main() {	
	logger := log.New(&buf, "Time: ", log.Lmicroseconds)
	logger.Println("Inicio")
	
	List = map[string]*Data{
		"8bitdeviants":{},
		"aikalee":{}, 
		"allysonbettie":{}, 
		"anya96":{},
		"bllueberrylove":{},
		"Caylin":{},
		"emmac":{},
		"evelynclaire":{},
		"fantasy_desires":{},
		"femmexfatale":{},
		"fiery_redhead":{},
		"gio_emilyouhot":{},
		"goldengoddessxxx":{},
		"ingennui":{},
		"lana_del_bae":{},
		//"lanadesex":{},
		"leticiablanee":{},
		"lilithlunatic":{},
		"livecleo":{},
		"lovinglexis":{}, 
		"macykennedy":{},
		"maryblack1":{}, 
		"marymoody":{},
		"mesmereyez":{},
		"miacamhot":{},
		"mike_chloe":{},
		"millyrobinson":{},
		"petitesage":{},
		"priscillawtff":{}, 
		"rileyraine420":{},
		"sadieheartsliam":{},
		"sasha_de_sade":{},
		"sexy_aymee":{},  
		"sissyashleigh":{},  
		"sophiecutegirl":{},
		"valerymont_":{},
		"witcher_":{},
		"Wynterbanx":{},
		"xemmarose":{},
	}
	
	IfBanned(List)
	Check(List)
	for X,Y:=range List{
		fmt.Println(X,Y)
	}
	logger.Println("Final")
	fmt.Print(&buf)
}

func Check(List map[string]*Data){
	for Model,User:=range List{
		User.Changed=false
		Old:=User.Status
		rsp, err := http.Get(MainPage+Model)
		if err != nil {
			log.Fatal(err) 
		}
		defer rsp.Body.Close()

		buf, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			log.Fatal(err) 
		}
		Status:=!strings.Contains(string(buf), Offline)
		switch (Status){
			case false:	
			case true:
		}
		if Status != Old{
			ActualizeStatus(Model, User, Status)
		}
	}
}

/*06/08/2017
toda la familia en un fin de semana, bastante interesante
*/

func ActualizeStatus(Model string, User *Data, Status bool){
	User.Changed=true
	User.Status=Status
	if Status == true{
		User.In=time.Now()	
		fmt.Println(Model,"is Online")	
	}else{
		User.Out=time.Now()
		fmt.Println(Model,"is Offline")
		if !User.In.IsZero(){ 
			fmt.Println(Model,"Conected by:", User.Out.Sub(User.In))
		}
	}
}

func IfBanned(List map[string]*Data){
	for Model:=range List{
		rsp, err := http.Get(MainPage+Model)
		if err != nil {
			log.Fatal(err) 
		}
		defer rsp.Body.Close()

		buf, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			log.Fatal(err) 
		}
		Status:=strings.Contains(string(buf), Banned)
		if Status{
			fmt.Println(Model,"is Banned, Deleting")
			delete(List,Model)
		}
	}
}