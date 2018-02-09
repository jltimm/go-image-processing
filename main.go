package main

import (
	"io/ioutil"
	"os"

	"github.com/jltimm/go-image-processing/convolutions"
	"github.com/jltimm/go-image-processing/utils"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) >= 2 {
		conv := args[0]
		fileName := args[1]

		data, err := ioutil.ReadFile(fileName)
		check(err)

		//TODO: reformatting, add ars to sobel. maybe add
		//method to get image data instead of just ReadFile
		if len(data) >= 0 && conv == "sobel" {
			utils.CheckIfFileExists()
			convolutions.Sobel()
		}
	}
}
