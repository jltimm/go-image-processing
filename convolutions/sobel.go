package convolutions

import (
	"fmt"
	"image"
	"image/color"

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

func getPixelValue(color color.Color) int8 {
	r, _, _, _ := color.RGBA()
	return int8(r)
}

// Does the actual math for calculating the gradients
func getGradients(img image.Gray, x int, y int) (int8, int8) {
	gx := (kernelX[2][2]*getPixelValue(img.At(x-1, y-1)) + (kernelX[2][1] * getPixelValue(img.At(x-1, y))) + (kernelX[2][0] * getPixelValue(img.At(x-1, y+1))) +
		kernelX[1][2]*getPixelValue(img.At(x, y-1)) + (kernelX[1][1] * getPixelValue(img.At(x, y))) + (kernelX[1][0] * getPixelValue(img.At(x, y+1))) +
		kernelX[0][2]*getPixelValue(img.At(x+1, y-1)) + (kernelX[0][1] * getPixelValue(img.At(x+1, y))) + (kernelX[0][0] * getPixelValue(img.At(x+1, y+1))))

	gy := (kernelY[2][2]*getPixelValue(img.At(x-1, y-1)) + (kernelY[2][1] * getPixelValue(img.At(x-1, y))) + (kernelY[2][0] * getPixelValue(img.At(x-1, y+1))) +
		kernelY[1][2]*getPixelValue(img.At(x, y-1)) + (kernelY[1][1] * getPixelValue(img.At(x, y))) + (kernelY[1][0] * getPixelValue(img.At(x, y+1))) +
		kernelY[0][2]*getPixelValue(img.At(x+1, y-1)) + (kernelY[0][1] * getPixelValue(img.At(x+1, y))) + (kernelY[0][0] * getPixelValue(img.At(x+1, y+1))))

	return gx, gy
}

// Loops through image, calculating sobel
func sobelOperator(img image.Gray) *image.Gray {
	var (
		bounds = img.Bounds()
		sobel  = image.NewGray(bounds)
	)
	for x := 1; x < bounds.Max.X-1; x++ {
		for y := 1; y < bounds.Max.Y-1; y++ {
			gx, gy := getGradients(img, x, y)
			fmt.Println(gx)
			fmt.Println(gy)
		}
	}
	return sobel
}

// Sobel applies sobel filter to an image
func Sobel(filename string) *image.Gray {
	if !utils.CheckIfFileExists(filename) {
		panic("The file does not exist")
	}

	img := utils.ConvertToGrayscaleFromFilename(filename)
	if img == nil {
		panic("img returned nil")
	}

	sobelImg := sobelOperator(*img)
	return sobelImg
}
