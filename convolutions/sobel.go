package convolutions

import (
	"image"
	"image/color"
	"math"

	"github.com/jltimm/go-image-processing/utils"
)

// Declare the kernels for the sobel operator
var (
	kernelX = [][]int8{
		{1, 0, -1},
		{2, 0, -2},
		{1, 0, -1},
	}
	kernelY = [][]int8{
		{1, 2, 1},
		{0, 0, 0},
		{-1, -2, -1},
	}
)

//TODO: write method that converts the image data into a 2D array to eliminate all the img.At calls
//TODO: rename convolutions to edge-detection
func getImageData(img image.NRGBA) [][]uint8 {
	bounds := img.Bounds()
	imgArray := make([][]uint8, bounds.Max.Y)
	for i := 0; i < bounds.Max.Y; i++ {
		imgArray[i] = make([]uint8, bounds.Max.X)
	}
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			imgArray[x][y] = uint8(r)
		}
	}
	return imgArray
}

// getPixelValue returns only the r value (since it's grayscale and they're all the same)
func getPixelValue(color color.Color) int8 {
	r, _, _, _ := color.RGBA()
	return int8(r)
}

// calculateGradients does the actual math for calculating the gradients
func calculateGradients(img image.NRGBA, x int, y int) (float64, float64, uint8) {
	//TODO: consider declaring all of the img.At so it's not found twice
	gx := (kernelX[2][2]*getPixelValue(img.At(x-1, y-1)) + (kernelX[2][1] * getPixelValue(img.At(x-1, y))) + (kernelX[2][0] * getPixelValue(img.At(x-1, y+1))) +
		kernelX[1][2]*getPixelValue(img.At(x, y-1)) + (kernelX[1][1] * getPixelValue(img.At(x, y))) + (kernelX[1][0] * getPixelValue(img.At(x, y+1))) +
		kernelX[0][2]*getPixelValue(img.At(x+1, y-1)) + (kernelX[0][1] * getPixelValue(img.At(x+1, y))) + (kernelX[0][0] * getPixelValue(img.At(x+1, y+1))))

	gy := (kernelY[2][2]*getPixelValue(img.At(x-1, y-1)) + (kernelY[2][1] * getPixelValue(img.At(x-1, y))) + (kernelY[2][0] * getPixelValue(img.At(x-1, y+1))) +
		kernelY[1][2]*getPixelValue(img.At(x, y-1)) + (kernelY[1][1] * getPixelValue(img.At(x, y))) + (kernelY[1][0] * getPixelValue(img.At(x, y+1))) +
		kernelY[0][2]*getPixelValue(img.At(x+1, y-1)) + (kernelY[0][1] * getPixelValue(img.At(x+1, y))) + (kernelY[0][0] * getPixelValue(img.At(x+1, y+1))))

	_, _, _, a := img.At(x, y).RGBA()

	return float64(gx), float64(gy), uint8(a)
}

// calculateMagniute calculates the magnitude
func calculateMagnitude(gx float64, gy float64) uint8 {
	g := math.Sqrt((gx * gx) + (gy * gy))
	if g > 255 {
		return 255
	}
	return uint8(g)
}

// Loops through image, calculating sobel
func sobelOperator(img image.NRGBA) *image.NRGBA {
	var (
		bounds = img.Bounds()
		sobel  = image.NewNRGBA(bounds)
	)
	for x := 1; x < bounds.Max.X-1; x++ {
		for y := 1; y < bounds.Max.Y-1; y++ {
			gx, gy, a := calculateGradients(img, x, y)
			g := calculateMagnitude(gx, gy)
			sobel.Set(x, y, color.RGBA{g, g, g, a})
		}
	}
	return sobel
}

// Sobel applies sobel filter to an image
func Sobel(filename string) *image.NRGBA {
	if !utils.CheckIfFileExists(filename) {
		panic("The file does not exist")
	}

	img := utils.ConvertToGrayscaleFromFilenameReturnNRGBA(filename)
	if img == nil {
		panic("img returned nil")
	}
	getImageData(*img)
	sobel := sobelOperator(*img)
	return sobel
}

// CreateSobelFromFile takes as input a filename, performs the sobel transform on the file, and
// creates a file with the name newFilename
func CreateSobelFromFile(filename string, newFilename string) {
	sobel := Sobel(filename)
	utils.CreateFileFromNRGBA(newFilename, sobel)
}
