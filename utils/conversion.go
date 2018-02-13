package utils

import (
	"image"
	"os"
)

// DecodeImage takes as input a filename, and returns the decoded Image data, and its bounds
// TODO: consider getting rid of ext := getFileExtension(filename)
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

// TODO: New method: ConvertToSepiaFromFilename / ConvertToSepiaFromImageData
// https://stackoverflow.com/questions/1061093/how-is-a-sepia-tone-created
