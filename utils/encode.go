package utils

import (
	"image"
	"image/png"
	"os"
)

//TODO: add methods to create files from simply a filename input

// CreateFileFromNRGBA takes as input image data and encodes it
// TODO: clean this up
func CreateFileFromNRGBA(filename string, newFilename string, img *image.NRGBA) {
	outputFile, err := os.Create("test-sobel.png")
	if err != nil {
		panic(err)
	}

	// TODO: split desiredFilename by ',' and check for extension. if no extension, use the one from the filename
	// Check if go can do optional parameters??

	png.Encode(outputFile, img)
	outputFile.Close()
}
