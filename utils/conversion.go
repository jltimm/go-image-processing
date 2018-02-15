package utils

import (
	"fmt"
	"image"
	"os"
)

// DecodeImage takes as input a filename, and returns the decoded Image data, and its bounds
// TODO: move this method to decode.go
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

// ConvertToGrayscaleFromImageData takes as input an image, and converts it to grayscale.
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

func convertTo8Bit(r uint32, g uint32, b uint32, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r), uint8(g), uint8(b), uint8(a)
}

// ConvertToSepiaFromImageData takes as input an image and converts it to sepia tone
func ConvertToSepiaFromImageData(img image.Image) *image.RGBA {
	var (
		bounds = img.Bounds()
		sepia  = image.NewRGBA(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var r, g, b, a = convertTo8Bit(img.At(x, y).RGBA())

			fmt.Println(r)
			fmt.Println(g)
			fmt.Println(b)
			fmt.Println(a)
		}
	}
	return sepia
}

// https://stackoverflow.com/questions/1061093/how-is-a-sepia-tone-created
// ConvertToSepiaFromFilename takes as input a filename and converts it to sepia tone
func ConvertToSepiaFromFilename(filename string) image.Image {
	img := DecodeImage(filename)
	return img
}
