package main

import (
	"log"
	"net/http"
)

//main - Main function
func main() {
	//Mux and handlers
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/loadall", loadAll)

	//Web server configuration
	serv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("Started on %s", serv.Addr)

	//Main application thread
	log.Fatal(serv.ListenAndServe())
}
