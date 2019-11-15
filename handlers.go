package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type packet struct {
	id         int
	timeRel    float32
	ipSrc      string
	ipDest     string
	portSrc    string
	portDest   string
	packetType string
}

//index - Handles main page
func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Template parsing error: %s", err)
		return
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

//index - Handles view page
func view(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Template parsing error: %s", err)
		return
	}
	tmpl.ExecuteTemplate(w, "view", nil)
}

//loadAll - Load packets list from db
func loadAll(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/packets.html")
	if err != nil {
		log.Fatalf("Template parsing error: %s", err)
		return
	}
	tmpl.Execute(w, nil)
}

//upload - File upload method
func upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Upload error: %s", err)
		return
	}
	defer file.Close()

	fileID := md5sum(header.Filename)
	fileName := header.Filename

	out, err := os.Create(fmt.Sprintf(".cache/%s", fileID))
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Error: %s", err)
		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintf(w, "Server io error: %s", err)
	}

	_, err = db.Exec("insert into uploads (file_id, file_name) values ($1, $2)", fileID, fileName)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/view", 301)
}
