package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "/home/murat-cileli/.local/share/mistory/database.db")
	catch(err)
	return db
}

func insertToDatabase(captureDateTime string, imagePath string, ocrText string) {
	statement, err := db.Prepare("INSERT INTO captures (capture_date_time, capture_file, ocr_text, ocr_bounds, capture_resolution, session_uuid) VALUES(?, ?, ?, ?, ?, ?);")
	catch(err)
	defer statement.Close()
	_, err = statement.Exec(1, imagePath, ocrText, nil, displayBounds.String(), sessionUuid)
	catch(err)
}
