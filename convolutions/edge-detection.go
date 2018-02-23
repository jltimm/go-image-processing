package convolutions

import (
	"image"
	"image/color"

	"github.com/jltimm/go-image-processing/utils"
)

// Declare the kernels for the sobel operator
var (
	sobelKernelX = [][]int8{
		{1, 0, -1},
		{2, 0, -2},
		{1, 0, -1},
	}
	sobelKernelY = [][]int8{
		{1, 2, 1},
		{0, 0, 0},
		{-1, -2, -1},
	}
)

// Loops through image, calculating sobel
func sobelOperator(img image.NRGBA) *image.NRGBA {
	var (
		bounds   = img.Bounds()
		sobel    = image.NewNRGBA(bounds)
		imgArray = utils.GetImageData(img)
	)
	for x := 1; x < bounds.Max.X-1; x++ {
		for y := 1; y < bounds.Max.Y-1; y++ {
			gx, gy, a := utils.CalculateGradients(imgArray, sobelKernelX, sobelKernelY, x, y)
			g := utils.CalculateMagnitude(gx, gy)
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
	sobel := sobelOperator(*img)
	return sobel
}

// CreateSobelFromFile takes as input a filename, performs the sobel transform on the file, and
// creates a file with the name newFilename
func CreateSobelFromFile(filename string, newFilename string) {
	sobel := Sobel(filename)
	utils.CreateFileFromNRGBA(newFilename, sobel)
}
