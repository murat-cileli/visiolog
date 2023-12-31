package main

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func createDatabaseIfNotExists() {
	if _, err := os.Stat(filepath.Join(appDataDir, "visiolog.db")); os.IsNotExist(err) {
		_, err := os.Create(filepath.Join(appDataDir, "visiolog.db"))
		catch(err)
		db, err := sql.Open("sqlite3", filepath.Join(appDataDir, "visiolog.db"))
		catch(err)
		defer db.Close()

		sqlStmt := `
			CREATE TABLE captures (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			capture_date_time TEXT NOT NULL,
			capture_file TEXT NOT NULL,
			capture_resolution TEXT NOT NULL,
			capture_interval INTEGER DEFAULT 0 NOT NULL,
			session_uuid TEXT NOT NULL,
			hocr_text TEXT
			);
		`
		_, err = db.Exec(sqlStmt)
		catch(err)
	}
}

func openDatabase() *sql.DB {
	createDatabaseIfNotExists()
	db, err := sql.Open("sqlite3", filepath.Join(appDataDir, "visiolog.db"))
	catch(err)
	return db
}

func insertToDatabase(captureDateTime string, captureFileName string, hOcrText string) {
	statement, err := db.Prepare("INSERT INTO captures (capture_date_time, capture_file, capture_resolution, capture_interval, session_uuid, hocr_text) VALUES(?, ?, ?, ?, ?, ?);")
	catch(err)
	defer statement.Close()
	_, err = statement.Exec(captureDateTime, captureFileName+".png", displayBounds.String(), captureInterval, sessionUuid, hOcrText)
	catch(err)
}
