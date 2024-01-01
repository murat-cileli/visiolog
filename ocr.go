package main

import (
	"bytes"
	"image/png"

	"github.com/otiai10/gosseract/v2"
)

type ocrType struct{}

func (ocr *ocrType) getHocrText() string {
	gosseractClient := gosseract.NewClient()
	gosseractClient.SetLanguage(captureOptions.ocrLanguages) // TODO
	defer gosseractClient.Close()

	buffer := new(bytes.Buffer)
	defer buffer.Reset()
	helper.catch(png.Encode(buffer, capture.image))
	helper.catch(gosseractClient.SetImageFromBytes(buffer.Bytes()))

	hOcrText, err := gosseractClient.HOCRText()
	helper.catch(err)

	return hOcrText
}
