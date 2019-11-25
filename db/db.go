package db

import "database/sql"

//DB - app variable for db access
var DB *sql.DB

//Init - init app database
func Init(name string) error {
	var err error
	//Open db
	DB, err = sql.Open("sqlite3", name)
	if err != nil {
		return err
	}
	return nil
}
