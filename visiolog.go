package main

import (
	"database/sql"
	"errors"
	"image"
	"time"

	"github.com/google/uuid"
	"github.com/kbinani/screenshot"
)

var displayBounds image.Rectangle = screenshot.GetDisplayBounds(0)
var sessionUuid string = uuid.New().String()
var appDataDir string
var db *sql.DB

const captureInterval = 3

func main() {
	if screenshot.NumActiveDisplays() <= 0 {
		catch(errors.New("No active display found."))
	}

	appDataDir = getAppDataDir()
	db = openDatabase()

	for {
		go capture()
		time.Sleep(captureInterval * time.Second)
	}
}
