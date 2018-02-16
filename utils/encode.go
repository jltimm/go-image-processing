package utils

import (
	"image"
	"image/png"
	"os"
)

//TODO: add methods to create files from simply a filename input

// CreateFileFromRGBA takes as input image data and encodes it
// TODO: clean this up
func CreateFileFromRGBA(filename string, desiredFilename string, img *image.RGBA) {
	outputFile, err := os.Create("test-gray.png")
	if err != nil {
		// TODO: handle
	}

	// TODO: split desiredFilename by ',' and check for extension. if no extension, use the one from the filename
	// Check if go can do optional parameters??

	png.Encode(outputFile, img)
	outputFile.Close()
}
