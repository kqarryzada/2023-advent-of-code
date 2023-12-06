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
	times := strings.Fields(fileLines[0])[1:]
	distances := strings.Fields(fileLines[1])[1:]

	winsPerRound := make([]int, 0)

	for round := 0; round < len(times); round++ {
		totalTime, _ := strconv.Atoi(times[round])
		distanceToBeat, _ := strconv.Atoi(distances[round])
		winsPerRound = append(winsPerRound, 0)

		for i := 1; i < totalTime; i++ {
			score := calculateDistance(i, totalTime)
			if score > distanceToBeat {
				winsPerRound[round]++
			}
		}
	}

	result := 1
	for _, roundVal := range winsPerRound {
		result *= roundVal
	}

	fmt.Printf("The product of the winning combinations is %d.\n", result)
}
