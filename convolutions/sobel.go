package convolutions

import (
	"fmt"

	"github.com/jltimm/go-image-processing/utils"
)

// Sobel applies sobel filter to an image
// TODO: write this function already!
func Sobel(filename string) {
	if !utils.CheckIfFileExists(filename) {
		fmt.Println("The file does not exist")
		return
	}
	img := utils.DecodeImage(filename)
	grayImg := utils.ConvertToGrayscaleFromImageData(img)
	fmt.Println(grayImg)
}
