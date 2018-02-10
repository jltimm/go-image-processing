package utils

import (
	"image"
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

// DecodeImage takes as input a filename, and returns the decoded Image data, and its bounds
func DecodeImage(filename string) (image.Image, image.Rectangle) {
	//TODO: consider adding check to see if this throws an error. it shouldn't, but who knows
	reader, _ := os.Open(filename)
	defer reader.Close()
	//TODO: check the file type of the image before decoding... this is hardcoded for PNG
	img, err := png.Decode(reader)
	if err != nil {
		panic(err)
	}
	return img, img.Bounds()
}
