package db

import "database/sql"

//DB - app variable for db access
var DB *sql.DB

//Init - init app database
func Init(name string) error {
	var err error
	//Open db
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}
