package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//main - Main function
func main() {
	//Open db
	db, err := sql.Open("sqlite3", "packets.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Mux and handlers
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/view", view)
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
