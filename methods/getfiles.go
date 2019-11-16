package methods

import (
	"log"
	"net/http"
)

//GetFiles - retrieve uploaded files data
func GetFiles(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from uploads")
	if err != nil {
		log.Printf("Error querying db: %s", err)
		return
	}
	defer func() {
		rows.Close()
		http.Redirect(w, r, "/", 301)
	}()

	uploads := []Upload{}

	for rows.Next() {
		u := Upload{}
		err := rows.Scan(&u.FileID, &u.FileName)
		if err != nil {
			log.Printf("Error scanning db response: %s", err)
			return
		}
		uploads = append(uploads, u)
	}
}
