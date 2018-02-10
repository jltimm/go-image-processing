package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// checkIfPngOrJpg takes as input an extension and returns true if its jpg or png.
// TODO: could do some refactoring on where this is called.
func checkIfPngOrJpg(ext string) bool {
	return (ext == "png" || ext == "jpg" || ext == "jpeg")
}

// getFileExtension takes as input a filename, and gets the extension
func getFileExtension(filename string) string {
	splitString := strings.Split(filename, ".")
	ext := strings.ToLower(splitString[len(splitString)-1])
	return ext
}

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
