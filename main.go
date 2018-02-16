package main

import (
	"os"

	"github.com/jltimm/go-image-processing/convolutions"
	"github.com/jltimm/go-image-processing/utils"
)

// Basically, all this method does is allow us to run
// the various tools in this repo.
func main() {
	args := os.Args[1:]
	if len(args) >= 2 {
		conv := args[0]
		filename := args[1]

		if conv == "sobel" {
			convolutions.Sobel(filename)
		}

		utils.CreateFileFromRGBA("", "", utils.ConvertToSepiaFromFilename(filename))
	}
}
