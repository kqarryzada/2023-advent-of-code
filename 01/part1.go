package main

import (
	"fmt"
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

	return fileAsSlice
}

func extractCalibratedValue(line string) (calibratedValue int) {
	// Assemble all numerical values in a slice.
	numbers := make([]int, 0)
	for _, char := range line {
		if char < '0' || char > '9' {
			continue
		}

		numericalValue := int(char) - '0'
		numbers = append(numbers, numericalValue)
	}

	// Fetch the first and final values in the slice.
	var first int
	var last int
	if len(numbers) == 0 {
		first = 0
		last = 0
	} else if len(numbers) == 1 {
		first = numbers[0]
		last = numbers[0]
	} else {
		first = numbers[0]
		last = numbers[len(numbers)-1]
	}

	return first*10 + last
}

func main() {
	sum := 0

	fileLines := LoadFile("input.txt")
	for _, line := range fileLines {
		value := extractCalibratedValue(line)
		sum += value
	}

	fmt.Printf("The total sum of the calibrated values is %d.\n", sum)
}
