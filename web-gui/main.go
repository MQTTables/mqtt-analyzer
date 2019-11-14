package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello TLS")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	serv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("Started on %s", serv.Addr)
	log.Fatal(serv.ListenAndServe())
}
