package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// CheckIfFileExists takes as input a filename, and checks if a file exists
func CheckIfFileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// decodePngOrJpg takes as input a file and extension, and returns the decoded image
func decodePngOrJpg(reader *os.File, ext string) image.Image {
	if ext == "png" {
		img, err := png.Decode(reader)
		if err != nil {
			panic(err)
		}
		return img
	}
	img, err := jpeg.Decode(reader)
	if err != nil {
		panic(err)
	}
	return img
}

// DecodeImage takes as input a filename, and returns the decoded Image data, and its bounds
func DecodeImage(filename string) image.Image {
	ext := getFileExtension(filename)
	if !checkIfPngOrJpg(ext) {
		panic("Unsupported file type: must be png or jpg")
	}
	reader, _ := os.Open(filename)
	defer reader.Close()
	img := decodePngOrJpg(reader, ext)
	return img
}
