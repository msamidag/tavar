package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", anasayfa)
	http.ListenAndServe(":8000", nil)
}
func anasayfa(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("anasayfa.html")
	if e != nil {
		log.Fatal(e)
	}
	t.Execute(w, nil)
}
