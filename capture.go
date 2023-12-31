package main

import (
	"errors"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/kbinani/screenshot"
)

func startCapture() {
	if screenshot.NumActiveDisplays() <= 0 {
		catch(errors.New("No active display found."))
	}

	for {
		go capture()
		time.Sleep(captureInterval * time.Second)
	}
}

func capture() {
	img, err := screenshot.Capture(0, 0, displayBounds.Dx(), displayBounds.Dy())
	catch(err)
	captureDateTime := time.Now()
	captureFileName := saveScreenshot(img, captureDateTime.Format("2006-01-02-15-04-05"))
	insertToDatabase(captureDateTime.Format("2006-01-02 15:04:05"), captureFileName, ocrFromScreenshot(img))
}

func saveScreenshot(img *image.RGBA, fileName string) string {
	imgPath := getCaptureSubDirsFromCaptureFileName(fileName)
	imgPath = filepath.Join(imgPath, fileName+".png")
	file, err := os.Create(imgPath)
	catch(err)
	defer file.Close()
	catch(png.Encode(file, img))

	return fileName
}
