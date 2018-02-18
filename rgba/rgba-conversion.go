package rgba

import (
	"image"
	"image/color"

	"github.com/jltimm/go-image-processing/utils"
)

// ConvertToGrayscaleFromImageData takes as input image data and returns a grayscale image
// TODO: These two methods are named horribly, and should be placed in a new folder. In fact, all conversion types with different
// return types should be moved to different folders so they can all use a common name.
// TODO: clean up
func ConvertToGrayscaleFromImageData(img image.Image) *image.RGBA {
	var (
		bounds = img.Bounds()
		gray   = image.NewRGBA(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var r, g, b, _ = utils.Convert32BitTo8Bit(img.At(x, y).RGBA())
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			pixel := color.Gray{uint8(lum)}
			gray.Set(x, y, pixel)
		}
	}
	return gray
}

// ConvertToGrayscaleFromFilename takes as input image data and returns a grayscale image
func ConvertToGrayscaleFromFilename(filename string) *image.RGBA {
	img := utils.DecodeImage(filename)
	return ConvertToGrayscaleFromImageData(img)
}
