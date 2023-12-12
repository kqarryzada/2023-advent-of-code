package utils

import (
	"log"
	"os"
	"strings"
)

// LoadFile obtains the entire contents of a file. This requires storing the
// full contents of the file in memory.
func LoadFile(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fileAsSlice := strings.Split(string(b), "\n")
	if fileAsSlice[len(fileAsSlice)-1] == "" {
		// Remove the last newline of the file from the slice, if it exists.
		fileAsSlice = fileAsSlice[:len(fileAsSlice)-1]
	}

	if len(fileAsSlice) == 0 {
		panic("The input file is empty.")
	}

	return fileAsSlice
}
