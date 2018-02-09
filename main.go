package main

import (
	"io/ioutil"
	"os"

	"github.com/jltimm/go-image-processing/convolutions"
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

		if len(data) >= 0 && conv == "sobel" {
			convolutions.Sobel()
		}
	}
}
