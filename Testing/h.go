package main
 
import (
"fmt"
"net/http"
)
 
func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	switch(r.FormValue("post")){
		case "456": 
			http.Redirect(w, r, "/Pato",200 )
			break
		default: 
			http.Redirect(w, r, "/Racoon", 301)
			break
	}

}
func Pato(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pato")
}

func Racoon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Racoon")
}

 
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", process)
	http.HandleFunc("/Pato", Pato)
	http.HandleFunc("/Racoon", Racoon)
	server.ListenAndServe()
}