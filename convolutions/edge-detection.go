package convolutions

import (
	"image"
	"image/color"

	"github.com/jltimm/go-image-processing/utils"
)

// TODO: Scharr - https://docs.opencv.org/3.3.0/d2/d2c/tutorial_sobel_derivatives.html

// Loops through image, calculating sobel
func kernelOperator(img image.NRGBA, kernelX [][]int8, kernelY [][]int8) *image.NRGBA {
	var (
		bounds   = img.Bounds()
		sobel    = image.NewNRGBA(bounds)
		imgArray = utils.GetImageData(img)
	)
	for x := 1; x < bounds.Max.X-1; x++ {
		for y := 1; y < bounds.Max.Y-1; y++ {
			gx, gy, a := utils.CalculateGradients(imgArray, kernelX, kernelY, x, y)
			g := utils.CalculateMagnitude(gx, gy)
			sobel.Set(x, y, color.RGBA{g, g, g, a})
		}
	}
	return sobel
}

func Scharr(filename string) *image.NRGBA {
	var (
		scharrKernelX = [][]int8{
			{-3, 0, 3},
			{-10, 0, 10},
			{-3, 0, 3},
		}
		scharrKernelY = [][]int8{
			{-3, -10, -3},
			{0, 0, 0},
			{3, 10, 3},
		}
	)
	if !utils.CheckIfFileExists(filename) {
		panic("The file does not exist")
	}
	img := utils.ConvertToGrayscaleFromFilenameReturnNRGBA(filename)
	if img == nil {
		panic("img returned nil")
	}
	scharr := kernelOperator(*img, scharrKernelX, scharrKernelY)
	return scharr
}

// TODO: fix this and sobel, move some of the repetition out
// Prewitt applies prewitt operator to an image
func Prewitt(filename string) *image.NRGBA {
	var (
		prewittKernelX = [][]int8{
			{-1, 0, 1},
			{-1, 0, 1},
			{-1, 0, 1},
		}
		prewittKernelY = [][]int8{
			{-1, -1, -1},
			{0, 0, 0},
			{1, 1, 1},
		}
	)
	if !utils.CheckIfFileExists(filename) {
		panic("The file does not exist")
	}
	img := utils.ConvertToGrayscaleFromFilenameReturnNRGBA(filename)
	if img == nil {
		panic("img returned nil")
	}
	prewitt := kernelOperator(*img, prewittKernelX, prewittKernelY)
	return prewitt
}

// Sobel applies sobel filter to an image
func Sobel(filename string) *image.NRGBA {
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
	if !utils.CheckIfFileExists(filename) {
		panic("The file does not exist")
	}
	img := utils.ConvertToGrayscaleFromFilenameReturnNRGBA(filename)
	if img == nil {
		panic("img returned nil")
	}
	sobel := kernelOperator(*img, sobelKernelX, sobelKernelY)
	return sobel
}

// CreatePrewittFromFile takes as input a filename, performs the roberts cross on the file, and
// creates a file with the name newFilename
func CreatePrewittFromFile(filename string, newFilename string) {
	prewitt := Prewitt(filename)
	utils.CreateFileFromNRGBA(newFilename, prewitt)
}

// CreateSobelFromFile takes as input a filename, performs the sobel transform on the file, and
// creates a file with the name newFilename
func CreateSobelFromFile(filename string, newFilename string) {
	sobel := Sobel(filename)
	utils.CreateFileFromNRGBA(newFilename, sobel)
}
