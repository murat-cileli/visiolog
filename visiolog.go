package main

import (
	"database/sql"
	"flag"
	"image"
	"os"

	"github.com/google/uuid"
	"github.com/kbinani/screenshot"
)

var displayBounds image.Rectangle = screenshot.GetDisplayBounds(0)
var sessionUuid string = uuid.New().String()
var appDataDir string = getAppDataDir()
var db *sql.DB = openDatabase()
var capture captureType
var captureOptions captureOptionsType
var gui App

func main() {
	if len(os.Args) == 1 {
		gui.start()
	} else {
		isCaptureFlagSet := flag.Bool("capture", false, "Start capture mode.")
		flag.IntVar(&captureOptions.interval, "interval", 5, "Capture interval in seconds.")
		flag.StringVar(&captureOptions.ocrLanguages, "ocr-languages", "eng", "OCR language(s). Multiple language codes can be specified, separated by comma. For list of language codes, see https://tesseract-ocr.github.io/tessdoc/Data-Files-in-different-versions.html")
		flag.Parse()
		if *isCaptureFlagSet {
			capture.start()
		} else {
			flag.PrintDefaults()
		}
	}

}
