package main

import (
	"bytes"
	"image"
	"image/png"

	"github.com/otiai10/gosseract/v2"
)

func ocrFromScreenshot(img *image.RGBA) string {
	client := gosseract.NewClient()
	client.SetLanguage("eng", "tur")
	defer client.Close()

	buffer := new(bytes.Buffer)
	defer buffer.Reset()
	catch(png.Encode(buffer, img))
	catch(client.SetImageFromBytes(buffer.Bytes()))

	text, err := client.Text()
	catch(err)

	return text
}
