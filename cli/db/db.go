package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

const dbPath = "/opt/kprobe/db/db.sqlite"
var DB *sql.DB

func InitDatabase() {
	var err error

	if DatabaseExist() {
		err = os.Remove(dbPath)
		if err != nil {
			log.Fatal("Failed to delete database:", err)
		}
	}

	log.Println("Creating new database:", dbPath)

	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS kv_store (
		key VARCHAR(16) PRIMARY KEY,
		value VARCHAR(512) NOT NULL
	);
	`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Database initialized successfully.")
}

func DatabaseExist() bool {
	if _, err := os.Stat(dbPath); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		log.Fatal("Failed to check database existence:", err)
	}

	return false
}
