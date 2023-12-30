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
var db *sql.DB = openDatabase()

func main() {
	if screenshot.NumActiveDisplays() <= 0 {
		catch(errors.New("No active display found."))
	}

	for {
		go capture()
		time.Sleep(5 * time.Second)
	}
}
