package utils

import (
	"image"
	"image/png"
	"os"
)

//TODO: add methods to create files from simply a filename input

// CreateFileFromRGBA takes as input image data and encodes it
// TODO: clean this up
func CreateFileFromRGBA(filename string, convolutionName string, img *image.RGBA) {
	outputFile, err := os.Create("test.png")
	if err != nil {
		// TODO: handle
	}

	png.Encode(outputFile, img)
	outputFile.Close()
}
