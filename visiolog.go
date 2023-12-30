package main

import (
	"database/sql"
	"errors"
	"fmt"
	"image"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kbinani/screenshot"
)

var displayBounds image.Rectangle = screenshot.GetDisplayBounds(0)
var sessionUuid string = uuid.New().String()
var appDataDir string = getAppDataDir()
var db *sql.DB = openDatabase()

const captureInterval = 3

func main() {
	switch os.Args[1] {
	case "capture":
		startCapture()
	case "gui":
		startGui()
	default:
		catch(errors.New("Invalid argument."))
	}
}

func startCapture() {
	if screenshot.NumActiveDisplays() <= 0 {
		catch(errors.New("No active display found."))
	}

	for {
		go capture()
		time.Sleep(captureInterval * time.Second)
	}
}

func startGui() {
	// TODO
	fmt.Println("Work in progress.")
}
