package methods

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//GetPackets - Load packets list html from db
func GetPackets(w http.ResponseWriter, r *http.Request) {
	packets := []Packet{}

	rows, err := db.Query(fmt.Sprintf("select * from '%s'", r.FormValue("fileid")))
	if err != nil {
		log.Printf("Error querying db: %s", err)
		return
	}
	defer func() {
		rows.Close()
		http.Redirect(w, r, "/", 301)
	}()

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
