package db

import (
	"database/sql"
	"errors"
	"os"
	"UptimeKumaProbeAPI/helpers"
	_ "modernc.org/sqlite"
)

const (
	DB_CONNECTION_FAILED   string = "Failed to connect to database"
	DB_QUERY_FAILED        string = "Failed to get value from database"
	DB_SCAN_NEWEST_FAILED  string = "Failed to get newest scan from database"
	RES_OK				   string = "OK"
)

// const dbPath = "/opt/kprobe/db/db.sqlite"
const dbPath = "../cli/db.sqlite" //FOR TESTING, CHANGE TO ABOVE

var DB *sql.DB

func connectDatabase() bool {
	if !DatabaseExist() {
		helpers.PrintError("Database does not exist, run CLI app with <kprobe db init> first")
		return false
	}

	var err error

	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		helpers.PrintError("Failed to connect to database")
		return false
	}

	return true
}

func DatabaseExist() bool {
	if _, err := os.Stat(dbPath); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		helpers.PrintError("Failed to check database existence (" + err.Error() + ")")
	}

	return false
}

func GetValue(key string) (string, string) {
	if DB == nil {
		if !connectDatabase() {
			return "", DB_CONNECTION_FAILED
		}
	}

	var value string

	query := `
	SELECT value
	FROM keys
	WHERE name = ?;
	`

	err := DB.QueryRow(query, key).Scan(&value)
	if err != nil {
		helpers.PrintError("Failed to get value from database (" + err.Error() + ")")
		return "", DB_QUERY_FAILED
	}

	return value, RES_OK
}

func GetScanNewest(scanName string) (helpers.ScanRes, string) {
	if DB == nil {
		if !connectDatabase() {
			return helpers.ScanRes{}, DB_CONNECTION_FAILED
		}
	}

	query := `
	SELECT generated, passed
	FROM history
	WHERE scan_name = ?
	ORDER BY generated DESC
	LIMIT 1;
	`

	var res helpers.ScanRes

	err := DB.QueryRow(query, scanName).Scan(&res.Generated, &res.Passed)
	if err != nil {
		if err == sql.ErrNoRows {
			return helpers.ScanRes{}, DB_SCAN_NEWEST_FAILED
		}
		helpers.PrintError("Failed to get data from database (" + err.Error() + ")")
		return helpers.ScanRes{}, DB_CONNECTION_FAILED
	}

	return res, RES_OK
}