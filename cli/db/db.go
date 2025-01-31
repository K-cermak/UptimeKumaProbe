package db

import (
	"database/sql"
	"errors"
	"os"
	"time"
	_ "modernc.org/sqlite"
	"UptimeKumaProbe/helpers"
)

//const dbPath = "/opt/kprobe/db/db.sqlite"
const dbPath = "db.sqlite" //TODO temporary, change

var DB *sql.DB

func InitDatabase() {
	var err error

	if DatabaseExist() {
		err = os.Remove(dbPath)
		if err != nil {
			helpers.PrintError(true, "Failed to delete database (" + err.Error() + ")")
		}
	}

	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		helpers.PrintError(true, "Failed to connect to database (" + err.Error() + ")")
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS keys (
		name VARCHAR(32) PRIMARY KEY,
		value VARCHAR(4096) NOT NULL
	);
	`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		helpers.PrintError(true, "Failed to create table (" + err.Error() + ")")
	}

	InsertDbValue("probe_name", "New Probe")
	InsertDbValue("db_version", "v1.0")
	InsertDbValue("db_init_time", time.Now().String()) 
	InsertDbValue("config_set", "false")
	InsertDbValue("api_port", "80")
	InsertDbValue("editor_endpoint", "true")
}

func DatabaseExist() bool {
	if _, err := os.Stat(dbPath); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		helpers.PrintError(true, "Failed to check database existence (" + err.Error() + ")")
	}

	return false
}

func InsertDbValue(key string, value string) {
	insertQuery := `
	INSERT INTO keys (name, value) 
	VALUES (?, ?) 
	ON CONFLICT(name) 
	DO UPDATE SET value = excluded.value;
	`

	_, err := DB.Exec(insertQuery, key, value)
	if err != nil {
		helpers.PrintError(true, "Failed to insert data into database (" + err.Error() + ")")
	}
}