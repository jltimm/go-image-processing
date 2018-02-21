package utils

import (
	"image"
	"image/png"
	"os"
)

//TODO: add methods to create files from simply a filename input

// CreateFileFromNRGBA takes as input image data and encodes it
// TODO: clean this up
func CreateFileFromNRGBA(filename string, img *image.NRGBA) {
	outputFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	// TODO: split desiredFilename by ',' and check for extension. if no extension, use the one from the filename
	png.Encode(outputFile, img)
	outputFile.Close()
}
