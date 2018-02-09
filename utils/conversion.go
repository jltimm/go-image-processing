package utils

import (
	"os"
)

// CheckIfFileExists takes as input a filename, and checks if a file exists
func CheckIfFileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}
