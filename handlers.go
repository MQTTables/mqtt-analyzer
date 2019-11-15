package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	uuid "github.com/satori/go.uuid"
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

type uploads struct {
	fileID   string
	fileName string
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
	row := db.QueryRow("select * from uploads")
	upl := uploads{}
	err := row.Scan(&upl.fileID, &upl.fileName)
	if err != nil {
		log.Printf("Error scanning db response: %s", err)
		return
	}

	cmd := exec.Command("python3", "p-modules/p_main.py", "packets", upl.fileID, "pcap")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Python module error: %s", err)
		return
	}
	log.Println(string(out))

	packets := []packet{}

	rows, err := db.Query(fmt.Sprintf("select * from '%s'", upl.fileID))
	if err != nil {
		log.Printf("Error querying db: %s", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := packet{}
		err := rows.Scan(&p.id, &p.timeRel, &p.ipSrc, &p.ipDest, &p.portSrc, &p.portDest, &p.packetType)
		if err != nil {
			log.Printf("Error scanning db response: %s", err)
			return
		}
		packets = append(packets, p)
	}

	tmpl, err := template.ParseFiles("templates/packets.html")
	if err != nil {
		log.Fatalf("Template parsing error: %s", err)
		return
	}
	tmpl.Execute(w, packets)
}

//upload - File upload method
func upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Upload error: %s", err)
		return
	}
	defer file.Close()

	fileID := uuid.Must(uuid.NewV4()).String()
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
