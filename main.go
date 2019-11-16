package main

import (
	"log"
	"mqtt-analyzer/db"
	"mqtt-analyzer/methods"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//main - Main function
func main() {
	//Init database
	if err := db.Init("packets.db"); err != nil {
		log.Fatal(err)
	}

	//Create uploads table
	_, err := db.DB.Exec(`create table if not exists uploads (
		file_id text,
		file_name text
	);`)
	if err != nil {
		log.Fatal(err)
	}

	//Mux and handlers
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", methods.Index)
	mux.HandleFunc("/view", methods.View)
	mux.HandleFunc("/upload", methods.Upload)
	mux.HandleFunc("/getpackets", methods.GetPackets)
	mux.HandleFunc("/getfiles", methods.GetFiles)

	//Web server configuration
	serv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("Started on %s", serv.Addr)

	//Main application thread
	log.Fatal(serv.ListenAndServe())
}
