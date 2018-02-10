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

// DecodeImage takes as input a filename, and returns the decoded Image data, and its bounds
func DecodeImage(filename string) image.Image {
	ext := getFileExtension(filename)
	if !checkIfPngOrJpg(ext) {
		panic("Unsupported file type: must be png or jpg")
	}
	//TODO: consider adding check to see if this throws an error. it shouldn't, but who knows
	reader, _ := os.Open(filename)
	defer reader.Close()
	img := decodePngOrJpg(reader, ext)
	return img
}

// ConvertToGrayscaleFromImageData takes as input an image and the images bounds, and converts it to grayscale.
// A lot of this code is inspired by https://maxhalford.github.io/blog/halftoning-1/
// TODO: maybe move this out of this file?
// TODO: test if this works correctly
// TODO: consider adding checks on extension, existence, etc
func ConvertToGrayscaleFromImageData(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}

// ConvertToGrayscaleFromFilename takes as input a filename and converts it to grayscale
func ConvertToGrayscaleFromFilename(filename string) *image.Gray {
	img := DecodeImage(filename)
	return ConvertToGrayscaleFromImageData(img)
}
