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

func loadAll(w http.ResponseWriter, r *http.Request) {

}

func upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Upload error: %s", err)
		return
	}
	defer file.Close()

	out, err := os.Create(header.Filename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Error: %s", err)
		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintf(w, "Server io error: %s", err)
	}
	http.Redirect(w, r, "/", 301)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/upload", upload)
	serv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("Started on %s", serv.Addr)
	log.Fatal(serv.ListenAndServe())
}
