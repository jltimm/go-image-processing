package convolutions

import (
	"fmt"
	"image"

	"github.com/jltimm/go-image-processing/utils"
)

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

func getGradients(img image.Gray, x int, y int) (int, int) {

	return x, y
}

func sobelOperator(img image.Gray) image.Gray {
	bounds := img.Bounds()
	for x := 1; x < bounds.Max.X-1; x++ {
		for y := 1; y < bounds.Max.Y-1; y++ {
			//r, _, _, _ := img.At(x, y).RGBA()
			gx, gy := getGradients(img, x, y)
			fmt.Println(gx)
			fmt.Println(gy)
		}
	}
	return img
}

// Sobel applies sobel filter to an image
// TODO: write this function already!
func Sobel(filename string) image.Gray {
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
