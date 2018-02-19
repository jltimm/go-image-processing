package convolutions

import (
	"image"

	"github.com/jltimm/go-image-processing/utils"
)

func sobelOperator(img image.RGBA) image.RGBA {
	return img
}

// Sobel applies sobel filter to an image
// TODO: write this function already!
func Sobel(filename string) image.RGBA {
	if !utils.CheckIfFileExists(filename) {
		panic("The file does not exist")
	}

	img := utils.ConvertToGrayscaleFromFilenameReturnRGBA(filename)
	if img == nil {
		// TODO: handle, there may be a better way to do this
	}
	sobelImg := sobelOperator(*img)
	return sobelImg
}
