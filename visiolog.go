package main

import (
	"database/sql"
	"errors"
	"image"
	"os"

	"github.com/google/uuid"
	"github.com/kbinani/screenshot"
)

var displayBounds image.Rectangle = screenshot.GetDisplayBounds(0)
var sessionUuid string = uuid.New().String()
var appDataDir string = getAppDataDir()
var db *sql.DB = openDatabase()

const captureInterval = 3

func main() {
	switch {
	case len(os.Args) == 1:
		startGui()
	case os.Args[1] == "capture":
		startCapture()
	default:
		catch(errors.New("Invalid argument."))
	}
}
