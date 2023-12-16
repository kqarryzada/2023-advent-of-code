package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
	"strings"
)

func sequenceHash(sequence string) int {
	hash := 0
	for _, char := range sequence {
		hash += int(char)
		hash *= 17
		hash %= 256
	}

	return hash
}

func main() {
	fileLines := utils.LoadFile("input.txt")
	if len(fileLines) != 1 {
		panic("Unexpected input file format.")
	}

	initializationSequence := strings.Split(fileLines[0], ",")
	sum := 0
	for _, sequence := range initializationSequence {
		sum += sequenceHash(sequence)
	}

	fmt.Printf("The sum of all hashes is %d.\n", sum)
}
