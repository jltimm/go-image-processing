package main

import (
	"os"

	"github.com/jltimm/go-image-processing/utils"
)

// Basically, all this method does is allow us to run
// the various tools in this repo.
// TODO: general: consider dereferencing all pointers
func main() {
	args := os.Args[1:]
	if len(args) >= 2 {
		conv := args[0]
		filename := args[1]

		if conv == "sobel" {
			//convolutions.CreateSobelFromFile(filename, "sobel.jpg")
			// sepia := utils.ConvertToSepiaFromFilename(filename)
			// utils.CreateFileFromNRGBA("sepia.jpg", sepia)
			img := utils.ConvertToGrayscaleFromFilenameReturnNRGBA(filename)
			utils.CreateFileFromNRGBA("test1.jpg", img)
		}

	}
}
