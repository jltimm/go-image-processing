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
	defer outputFile.Close()
	if err != nil {
		panic(err)
	}
	ext := getFileExtension(filename)
	if !checkIfPngOrJpg(ext) {
		panic("File is not png or jpeg")
	}
	//TODO: figure out why jpeg.Encode is so horrible
	png.Encode(outputFile, img)
}
