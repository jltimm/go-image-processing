package utils

import (
	"image"
	"image/png"
	"os"
)

// CreateFileFromRGBA takes as input image data and encodes it
func CreateFileFromRGBA(filename string, convolutionName string, img *image.RGBA) {
	outputFile, err := os.Create("test.png")
	if err != nil {
		// TODO: handle
	}

	png.Encode(outputFile, img)
	outputFile.Close()
}
