package utils

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"strings"
)

// CalculateMagnitude calculates the magnitude
// TODO: replace 255 with a value from image
func CalculateMagnitude(gx float64, gy float64) uint8 {
	g := math.Sqrt((gx * gx) + (gy * gy))
	if g > 255 {
		return 255
	}
	return uint8(g)
}

// CalculateGradients does the actual math for calculating the gradients
// TODO: replace with for loops, so its not hardcoded
func CalculateGradients(imgArray [][]int8, kernelX [][]int8, kernelY [][]int8, x int, y int) (float64, float64, uint8) {
	// gx := (kernelX[2][2] * imgArray[x-1][y-1]) + (kernelX[2][1] * imgArray[x-1][y]) + (kernelX[2][0] * imgArray[x-1][y+1]) +
	// 	(kernelX[1][2] * imgArray[x][y-1]) + (kernelX[1][1] * imgArray[x][y]) + (kernelX[1][0] * imgArray[x][y+1]) +
	// 	(kernelX[0][2] * imgArray[x+1][y-1]) + (kernelX[0][1] * imgArray[x+1][y]) + (kernelX[0][0] * imgArray[x+1][y+1])

	// gy := (kernelY[2][2] * imgArray[x-1][y-1]) + (kernelY[2][1] * imgArray[x-1][y]) + (kernelY[2][0] * imgArray[x-1][y+1]) +
	// 	(kernelY[1][2] * imgArray[x][y-1]) + (kernelY[1][1] * imgArray[x][y]) + (kernelY[1][0] * imgArray[x][y+1]) +
	// 	(kernelY[0][2] * imgArray[x+1][y-1]) + (kernelY[0][1] * imgArray[x+1][y]) + (kernelY[0][0] * imgArray[x+1][y+1])
	// TODO: multiple var declaration
	gx := 0.0
	gy := 0.0
	xLength := len(kernelX[0])
	yLength := len(kernelY[0])
	fmt.Println(xLength)
	fmt.Println(yLength)

	return float64(gx), float64(gy), 255
}

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
