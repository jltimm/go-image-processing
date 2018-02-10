package utils

import (
	"image"
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

// DecodeImage takes as input a filename, and returns the decoded Image data, and its bounds
func DecodeImage(filename string) (image.Image, image.Rectangle) {
	ext := getFileExtension(filename)
	if !checkIfPngOrJpg(ext) {
		panic("Unsupported file type: must be png or jpg")
	}
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

// ConvertToGrayscale takes as input an image and the images bounds, and converts it to grayscale.
// A lot of this code is inspired by https://maxhalford.github.io/blog/halftoning-1/
// TODO: maybe move this out of this file?
// TODO: test if this works correctly
func ConvertToGrayscale(img image.Image, bounds image.Rectangle) *image.Gray {
	gray := image.NewGray(bounds)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}
