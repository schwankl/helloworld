package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] != "navarkos" {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	} else if r.URL.Path[1:] == "Kire" {
		fmt.Fprintf(w, "Hi there, Kire what you had for breakfast!")
	} else {
		fmt.Fprintf(w, "Hi there, Navarkos is the best thing that ever happened to kubernetes community!")
	}
}

func main() {
	log.Printf("Starting http server on port 8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
