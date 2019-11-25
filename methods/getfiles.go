package methods

import (
	"html/template"
	"log"
	"mqtt-analyzer/db"
	"net/http"
)

//FileData - struct of uploaded data
type FileData struct {
	FileID   string `json:"fileid"`
	FileName string `json:"filename"`
}

//GetFiles - retrieve uploaded files data
func GetFiles(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("select * from uploads")
	if err != nil {
		log.Printf("Error querying db: %s", err)
		return
	}
	defer rows.Close()

	uploads := []FileData{}

	for rows.Next() {
		u := FileData{}
		err := rows.Scan(&u.FileID, &u.FileName)
		if err != nil {
			log.Printf("Error scanning db response: %s", err)
			return
		}
		log.Println(u)
		uploads = append(uploads, u)
	}

	log.Println(uploads)

	tmpl, err := template.ParseFiles("templates/files.html")
	if err != nil {
		log.Fatalf("Template parsing error: %s", err)
		return
	}
	tmpl.Execute(w, uploads)
}
