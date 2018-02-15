package utils

import (
	"fmt"
	"image"
	"image/color"
)

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

func convert32BitTo8Bit(r, g, b, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r), uint8(g), uint8(b), uint8(a)
}

func getColor(r, g, b, a float64) color.Color {
	if r > a {
		r = a
	}
	if g > a {
		g = a
	}
	if b > a {
		b = a
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

// ConvertToSepiaFromImageData takes as input an image and converts it to sepia tone
func ConvertToSepiaFromImageData(img image.Image) *image.RGBA {
	var (
		bounds = img.Bounds()
		sepia  = image.NewRGBA(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var r, g, b, a = convert32BitTo8Bit(img.At(x, y).RGBA())
			rSepia := (float64(r) * .393) + (float64(g) * .769) + (float64(b) * .189)
			gSepia := (float64(r) * .349) + (float64(g) * .686) + (float64(b) * .168)
			bSepia := (float64(r) * .272) + (float64(g) * .534) + (float64(b) * .131)
			color := getColor(rSepia, gSepia, bSepia, float64(a))
			fmt.Println(color)
		}
	}
	return sepia
}

// ConvertToSepiaFromFilename takes as input a filename and converts it to sepia tone
func ConvertToSepiaFromFilename(filename string) image.Image {
	img := DecodeImage(filename)
	return img
}
