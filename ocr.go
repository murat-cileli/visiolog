package main

import (
	"bytes"
	"image"
	"image/png"

	"github.com/otiai10/gosseract/v2"
)

type ocrType struct{}

func (ocr *ocrType) getHocrText(img *image.RGBA) string {
	gosseractClient := gosseract.NewClient()
	gosseractClient.SetLanguage(captureOptions.ocrLanguages) // TODO
	defer gosseractClient.Close()

	buffer := new(bytes.Buffer)
	defer buffer.Reset()
	catch(png.Encode(buffer, img))
	catch(gosseractClient.SetImageFromBytes(buffer.Bytes()))

	hOcrText, err := gosseractClient.HOCRText()
	catch(err)

	return hOcrText
}
