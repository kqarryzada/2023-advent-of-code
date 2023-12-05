package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
	"slices"
	"strings"
)

func calculateScore(winCount int) int {
	if winCount == 0 {
		return 0
	}

	// The score value is 2^(n - 1) when at least one winning number is present.
	return 1 << (winCount - 1)
}

func processRound(line string) int {
	game := strings.Split(line, ": ")[1]
	gameData := strings.Split(game, "|")

	cardNumbers := strings.Split(gameData[0], " ")
	winningNumbers := strings.Split(gameData[1], " ")

	winningNumberCount := 0
	for _, number := range cardNumbers {
		if len(number) != 0 && slices.Contains(winningNumbers, number) {
			winningNumberCount++
		}
	}

	return calculateScore(winningNumberCount)
}

func main() {
	sum := 0
	fileLines := fileutils.LoadFile("input.txt")
	for _, line := range fileLines {
		sum += processRound(line)
	}

	fmt.Printf("The total number of points on the scratchcards is %d.\n", sum)
}
