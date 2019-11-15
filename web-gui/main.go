package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Template parsing error: %s", err)
	}
	tmpl.Execute(w, nil)
}

func loadInfo(w http.ResponseWriter, r *http.Request) {

}

func loadPackage(w http.ResponseWriter, r *http.Request) {

}

func upload(w http.ResponseWriter, r *http.Request) {
	upload, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "File upload unsuccsessful: %s", err)
		return
	}
	defer upload.Close()

	file, err := os.Create(header.Filename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. %s", err)
		return
	}

	defer file.Close()
	_, err = io.Copy(file, file)
	if err != nil {
		fmt.Fprintf(w, "Server file write error %s", err)
		return
	}
}

func readParams(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", index)
	serv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("Started on %s", serv.Addr)
	log.Fatal(serv.ListenAndServe())
}
