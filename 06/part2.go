package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
	"strconv"
	"strings"
)

// calculateDistance computes the distance that a boat will travel given the
// amount of time it is charged.
func calculateDistance(chargeTime int, totalTime int) int {
	timeToMove := totalTime - chargeTime
	return chargeTime * timeToMove
}

func main() {
	fileLines := fileutils.LoadFile("input.txt")
	timeString := strings.Join(strings.Fields(fileLines[0])[1:], "")
	distanceString := strings.Join(strings.Fields(fileLines[1])[1:], "")
	totalTime, _ := strconv.Atoi(timeString)
	distanceToBeat, _ := strconv.Atoi(distanceString)

	result := 0
	for i := 1; i < totalTime; i++ {
		score := calculateDistance(i, totalTime)
		if score > distanceToBeat {
			result++
		}
	}

	fmt.Printf("The number of winning combinations is %d.\n", result)
}
