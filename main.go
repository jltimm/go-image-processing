package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 2 {
		conv := args[0]
		fileName := args[1]
		fmt.Println(conv)
		fmt.Println(fileName)
		//TODO: search for file
	}
}
