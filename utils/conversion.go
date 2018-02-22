package utils

import (
	"image"
	"image/color"
)

// ConvertToGrayscaleFromImageData takes as input an image, and converts it to grayscale.
// A lot of this code is inspired by https://maxhalford.github.io/blog/halftoning-1/
// TODO: maybe move this out of this file?
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

// ConvertToSepiaFromImageData takes as input an image and converts it to sepia tone
// TODO: parameter to make more kinds of sepia
func ConvertToSepiaFromImageData(img image.Image) *image.NRGBA {
	var (
		bounds = img.Bounds()
		sepia  = image.NewNRGBA(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var r, g, b, a = convert32BitTo8Bit(img.At(x, y).RGBA())
			rSepia := (float64(r) * .393) + (float64(g) * .769) + (float64(b) * .189)
			gSepia := (float64(r) * .349) + (float64(g) * .686) + (float64(b) * .168)
			bSepia := (float64(r) * .272) + (float64(g) * .534) + (float64(b) * .131)
			color := getColor(rSepia, gSepia, bSepia, float64(a))
			sepia.Set(x, y, color)
		}
	}
	return sepia
}

// ConvertToSepiaFromFilename takes as input a filename and converts it to sepia tone
func ConvertToSepiaFromFilename(filename string) *image.NRGBA {
	img := DecodeImage(filename)
	return ConvertToSepiaFromImageData(img)
}

// ConvertToGrayscaleFromImageDataReturnNRGBA takes as input image data and returns a grayscale image
// TODO: These two methods are named horribly, and should be placed in a new folder. In fact, all conversion types with different
// return types should be moved to different folders so they can all use a common name.
// TODO: clean up
func ConvertToGrayscaleFromImageDataReturnNRGBA(img image.Image) *image.NRGBA {
	var (
		bounds = img.Bounds()
		gray   = image.NewNRGBA(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var r, g, b, _ = convert32BitTo8Bit(img.At(x, y).RGBA())
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			pixel := color.Gray{uint8(lum)}
			gray.Set(x, y, pixel)
		}
	}
	return gray
}

// ConvertToGrayscaleFromFilenameReturnNRGBA takes as input image data and returns a grayscale image
func ConvertToGrayscaleFromFilenameReturnNRGBA(filename string) *image.NRGBA {
	img := DecodeImage(filename)
	return ConvertToGrayscaleFromImageDataReturnNRGBA(img)
}
