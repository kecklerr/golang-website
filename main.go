package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello Ray!</h>")
}

func aboutPage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p>Here is the about  page.</p>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/about", aboutPage)
	http.ListenAndServe(":8080", r)
}
