package convolutions

import (
	"fmt"
	"image"

	"github.com/jltimm/go-image-processing/utils"
)

func sobelOperator(img image.Gray) image.Gray {
	bounds := img.Bounds()
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			fmt.Println(img.At(x, y))
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
		// TODO: handle, there may be a better way to do this
	}
	sobelImg := sobelOperator(*img)
	return sobelImg
}
