package main

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type databaseType struct {
	connection *sql.DB
	filePath   string
}

func (database *databaseType) createDatabaseIfNotExists() {
	if _, err := os.Stat(database.filePath); os.IsNotExist(err) {
		_, err := os.Create(database.filePath)
		helper.catch(err)
		database.connection, err = sql.Open("sqlite3", database.filePath)
		helper.catch(err)
		defer database.connection.Close()

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
		_, err = database.connection.Exec(sqlStmt)
		helper.catch(err)
	}
}

func (database *databaseType) initDatabase() {
	database.filePath = filepath.Join(helper.appDataDir, "visiolog.db")
	database.createDatabaseIfNotExists()
	var err error
	database.connection, err = sql.Open("sqlite3", database.filePath)
	helper.catch(err)
}
