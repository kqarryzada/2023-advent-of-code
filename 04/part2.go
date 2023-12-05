package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
	"slices"
	"strings"
)

// A 1-indexed slice that tracks the number of scratchcards. Index 0 is always
// set to a value of 0.
var scratchcards []int

func processRound(line string, roundNumber int) {
	game := strings.Split(line, ": ")[1]
	gameData := strings.Split(game, "|")

	cardNumbers := strings.Split(gameData[0], " ")
	winningNumbers := strings.Split(gameData[1], " ")

	newScratchCards := 0
	for _, number := range cardNumbers {
		if len(number) != 0 && slices.Contains(winningNumbers, number) {
			newScratchCards++
		}
	}

	for count := 0; count < newScratchCards; count++ {
		i := (roundNumber + 1) + count
		if i >= len(scratchcards) {
			// The index has exceeded the maximum scorecard number, so there is
			// nothing left to increment.
			break
		}

		// Increment by the number of copies that we currently have. For
		// example, if we have three scratchcards for round 5 and have two
		// winning numbers, then we will get three extra scratchcards for round
		// 6 and round 7.
		scratchcards[i] += scratchcards[roundNumber]
	}
}

func main() {
	fileLines := fileutils.LoadFile("input.txt")

	// Assemble the initial scratchcards array. We begin with 1 copy of every
	// scratchcard/game round (except "game 0"), so initialize all real values
	// to 1.
	scratchcards = make([]int, len(fileLines)+1)
	for i := range scratchcards {
		scratchcards[i] = 1
	}
	scratchcards[0] = 0

	for i, line := range fileLines {
		processRound(line, i+1)
	}

	sum := 0
	for _, val := range scratchcards {
		sum += val
	}

	fmt.Printf("The total number of scratchcards collected is %d.\n", sum)
}
