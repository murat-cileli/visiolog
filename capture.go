package main

import (
	"errors"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/kbinani/screenshot"
)

var ocr ocrType

type metaType struct {
	dateTime    time.Time
	fileName    string
	hOcrText    string
	ocrText     string
	resolution  string
	sessionUuid string
}

type captureType struct {
	image  *image.RGBA
	width  int
	height int
	meta   metaType
}

type captureOptionsType struct {
	interval     uint
	ocrLanguages string
}

func (capture *captureType) start() {
	if screenshot.NumActiveDisplays() <= 0 {
		catch(errors.New("No active display found."))
	}

	displayBounds := screenshot.GetDisplayBounds(0)
	capture.width = displayBounds.Dx()
	capture.height = displayBounds.Dy()

	capture.meta.resolution = strconv.Itoa(capture.width) + "x" + strconv.Itoa(capture.height)
	capture.meta.sessionUuid = uuid.New().String()

	intervalDuration := time.Duration(captureOptions.interval) * time.Second

	for {
		go capture.capture()
		time.Sleep(intervalDuration)
	}
}

func (capture *captureType) capture() {
	var err error
	capture.image, err = screenshot.Capture(0, 0, capture.width, capture.height)
	catch(err)
	capture.meta.dateTime = time.Now()
	capture.meta.hOcrText = ocr.getHocrText(capture.image)
	capture.saveToFile()
	capture.saveToDatabase()
}

func (capture *captureType) saveToFile() {
	capture.meta.fileName = capture.meta.dateTime.Format("2006-01-02-15-04-05") + ".png"
	imgPath := getCaptureSubDirsFromCaptureFileName(capture.meta.fileName)
	imgPath = filepath.Join(imgPath, capture.meta.fileName)
	file, err := os.Create(imgPath)
	catch(err)
	defer file.Close()
	catch(png.Encode(file, capture.image))
}

func (capture *captureType) saveToDatabase() {
	statement, err := db.Prepare("INSERT INTO captures (capture_date_time, capture_file, capture_resolution, capture_interval, session_uuid, hocr_text) VALUES(?, ?, ?, ?, ?, ?);")
	catch(err)
	defer statement.Close()
	_, err = statement.Exec(capture.meta.dateTime, capture.meta.fileName, capture.meta.resolution, captureOptions.interval, capture.meta.sessionUuid, capture.meta.hOcrText)
	catch(err)
}
