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

	_, err = db.Exec(`create table if not exists uploads (
		file_id text,
		file_name text
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into uploads (file_id, file_name) values ('test', 'n1')")
	if err != nil {
		panic(err)
	}

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
