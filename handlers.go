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

//Packet - basic struct of captured mqtt package
type Packet struct {
	ID         int
	TimeRel    float32
	IPSrc      string
	IPDest     string
	PortSrc    string
	PortDest   string
	PacketType string
}

//Upload - struct of uploaded data
type Upload struct {
	FileID   string
	FileName string
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
	upl := Upload{}
	err := row.Scan(&upl.FileID, &upl.FileName)
	if err != nil {
		log.Printf("Error scanning db response: %s", err)
		return
	}

	packets := []Packet{}

	rows, err := db.Query(fmt.Sprintf("select * from '%s'", upl.FileID))
	if err != nil {
		log.Printf("Error querying db: %s", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := Packet{}
		err := rows.Scan(&p.ID, &p.TimeRel, &p.IPSrc, &p.IPDest, &p.PortSrc, &p.PortDest, &p.PacketType)
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

	cmd := exec.Command("python3", "p-modules/p_main.py", "packets", fileID, "pcap")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Python module error: %s", err)
		return
	}

	http.Redirect(w, r, "/view", 301)
}
