package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
	"strings"
)

// An element represents a three-letter location with two neighboring
// directions.
type element struct {
	left  string
	right string

	// Indicates whether the fields end in the 'Z' character. This value is
	// saved so that each element computes it once.
	isLeftStop  bool
	isRightStop bool
}

// This parse function takes a line from the input file and converts it into an
// element object.
func parse(line string, pathMap *map[string]element) {
	fields := strings.Fields(line)

	newValue := new(element)
	location := fields[0]

	lastChar := len(fields[2]) - 1
	newValue.left = fields[2][1:lastChar]
	newValue.right = fields[3][0 : lastChar-1]

	newValue.isLeftStop = strings.HasSuffix(newValue.left, "Z")
	newValue.isRightStop = strings.HasSuffix(newValue.right, "Z")

	(*pathMap)[location] = *newValue
}

func getStartingNodes(fileLines []string) *[]string {
	returnNodes := make([]string, 0)
	for _, line := range fileLines {
		location := line[0:3]
		if strings.HasSuffix(location, "A") {
			returnNodes = append(returnNodes, location)
		}
	}

	return &returnNodes
}

// The input file for this problem is constructed in such a way that each path
// goes in a loop. To find the intersection point where all of these line up,
// this program finds the length of each loop length and computes the least
// common multiple, as determining this through brute force is not feasible.
func main() {
	fileLines := utils.LoadFile("input.txt")
	instructions := fileLines[0]

	pathMap := make(map[string]element, 0)

	for i := 2; i < len(fileLines); i++ {
		line := fileLines[i]
		parse(line, &pathMap)
	}

	// This array tracks the number of iterations it takes to reach an "ending"
	// element for each parallel path.
	pathIterationCounts := make([]int, 0)

	simultaneousPaths := getStartingNodes(fileLines[2:])
	for _, elem := range *simultaneousPaths {
		currentElement := elem
		numSteps := 0
		for i := 0; i < len(instructions); {
			numSteps++

			direction := instructions[i]
			var goLeft bool
			if direction == 'L' {
				goLeft = true
			} else if direction == 'R' {
				goLeft = false
			} else {
				panic("Invalid input: " + string(direction))
			}

			nextElement := pathMap[currentElement]
			var nextValue string
			var isStop bool
			if goLeft {
				nextValue = nextElement.left
				isStop = nextElement.isLeftStop
			} else {
				nextValue = nextElement.right
				isStop = nextElement.isRightStop
			}

			if isStop {
				pathIterationCounts = append(pathIterationCounts, numSteps)
				break
			}

			currentElement = nextValue
			if i == (len(instructions) - 1) {
				// When the instruction set runs out, start over from the
				// beginning.
				i = 0
			} else {
				i++
			}
		}
	}

	answer := utils.FindLCM(pathIterationCounts)
	fmt.Printf("The total number of steps required is %d.\n", answer)
}
