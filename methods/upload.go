package methods

import (
	"fmt"
	"io"
	"mqtt-analyzer/db"
	"net/http"
	"os"
	"os/exec"

	uuid "github.com/satori/go.uuid"
)

//Upload - File upload method
func Upload(w http.ResponseWriter, r *http.Request) {
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

	_, err = db.DB.Exec("insert into uploads (file_id, file_name) values ($1, $2)", fileID, fileName)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("python3", "p-modules/p_main.py", "packets", fileID, "pcap")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Python module error: %s", err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf(""), 301)
}
