package utils

import (
	"image"
	"image/color"
	"strings"
)

// GetImageData converts image data to a 2D array
func GetImageData(img image.NRGBA) [][]int8 {
	bounds := img.Bounds()
	imgArray := make([][]int8, bounds.Max.Y)
	for i := 0; i < bounds.Max.Y; i++ {
		imgArray[i] = make([]int8, bounds.Max.X)
	}
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			imgArray[x][y] = int8(r)
		}
	}
	return imgArray
}

// convert32BitTo8Bit converts uint32 to uint8
func convert32BitTo8Bit(r, g, b, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r), uint8(g), uint8(b), uint8(a)
}

// getColor takes as input rgba values, and if any of the values are greater
// than a, it just sets it to a
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
	return color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

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
