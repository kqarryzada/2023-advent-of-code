package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
	"strings"
)

type element struct {
	left  string
	right string
}

func parse(line string, pathMap *map[string]element) {
	fields := strings.Fields(line)

	newValue := new(element)
	location := fields[0]

	lastChar := len(fields[2]) - 1
	newValue.left = fields[2][1:lastChar]
	newValue.right = fields[3][0 : lastChar-1]

	(*pathMap)[location] = *newValue
}

func main() {
	fileLines := fileutils.LoadFile("input.txt")
	instructions := fileLines[0]

	pathMap := make(map[string]element, 0)

	for i := 2; i < len(fileLines); i++ {
		line := fileLines[i]
		parse(line, &pathMap)
	}

	instructionLength := len(instructions)
	currentElement := strings.Fields("AAA")[0]
	numSteps := 0
	for i := 0; i < instructionLength; {
		numSteps++

		direction := instructions[i]
		nextVal := pathMap[currentElement]

		var nextStep string
		if direction == 'L' {
			nextStep = nextVal.left
		} else if direction == 'R' {
			nextStep = nextVal.right
		} else {
			panic("Invalid input: " + string(direction))
		}

		currentElement = nextStep
		if currentElement == "ZZZ" {
			break
		}

		if i == (instructionLength - 1) {
			i = 0
		} else {
			i++
		}
	}

	fmt.Printf("Reached the 'ZZZ' step in %d steps.\n", numSteps)
}
