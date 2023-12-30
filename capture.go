package main

import (
	"image"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func capture() {
	img, err := screenshot.Capture(0, 0, displayBounds.Dx(), displayBounds.Dy())
	catch(err)
	captureDateTime := time.Now().Format("2006-01-02-15-04-05")
	imagePath := saveScreenshot(img, captureDateTime)
	insertToDatabase(captureDateTime, imagePath, ocrFromScreenshot(img))
}

func saveScreenshot(img *image.RGBA, fileName string) string {
	imgPath := "fs/" + fileName + ".png"
	file, err := os.Create(imgPath)
	catch(err)
	defer file.Close()
	catch(png.Encode(file, img))

	return imgPath
}
