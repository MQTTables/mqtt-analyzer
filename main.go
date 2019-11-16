package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//main - Main function
func main() {
	var err error
	//Open db
	db, err = sql.Open("sqlite3", "packets.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Create uploads table
	_, err = db.Exec(`create table if not exists uploads (
		file_id text,
		file_name text
	);`)
	if err != nil {
		log.Fatal(err)
	}

	//Mux and handlers
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/view", view)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/getpackets", getPackets)
	mux.HandleFunc("/getfiles", getFiles)

	//Web server configuration
	serv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("Started on %s", serv.Addr)

	//Main application thread
	log.Fatal(serv.ListenAndServe())
}
