package main

import (
    "net/http"
    "html/template"
	"log"
)


type User struct {
		Id int
		Name []string
		X int
	}
	
type UserList []User
var myuserlist User  = User {1, []string{"a", "b", "c"},19}
	

func main() {
	log.Println("Start")
	http.HandleFunc("/", Mapping)
	log.Println("Handling")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Mapping(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("Data.html") 
    if err != nil {
  	  log.Print("template parsing error: ", err) 
  	}
    err = t.Execute(w, myuserlist) 
    if err != nil { 
  	  log.Print("template executing error: ", err) 
  	}
	log.Println("Parsed")	
}