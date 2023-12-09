package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
	"sort"
	"strconv"
	"strings"
)

type pokerHand struct {
	// The raw string representing the poker hand.
	raw string

	// The value of the poker hand (e.g., a Full House)
	value pokerHandValue

	// The monetary bid associated with the poker hand.
	bid int
}

type pokerHandValue int

const (
	highCard pokerHandValue = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

// containsValue returns true if the input map has the specified value stored in
// the map, regardless of the key it is associated with.
func containsValue(m map[byte]int, v int) bool {
	for _, x := range m {
		if x == v {
			return true
		}
	}

	return false
}

// computeWildcard takes an existing poker hand and returns the value of that
// hand when wildcards are considered. For example, for a hand of four Aces and
// a wildcard Joker, the value becomes a five-of-a-kind.
func computeWildcard(currentHand pokerHandValue, numberWildcards int) pokerHandValue {
	var returnHand pokerHandValue
	if numberWildcards == 0 {
		return currentHand
	}
	if numberWildcards == 5 {
		return fiveOfAKind
	}

	if numberWildcards == 1 {
		switch currentHand {
		case highCard:
			returnHand = onePair
		case onePair:
			returnHand = threeOfAKind
		case twoPair:
			returnHand = fullHouse
		case threeOfAKind:
			returnHand = fourOfAKind
		case fourOfAKind:
			returnHand = fiveOfAKind
		default:
			panic(fmt.Sprintf("Unexpected hand type with %d wildcard: %d", numberWildcards, currentHand))
		}

		return returnHand
	}

	if numberWildcards == 2 {
		switch currentHand {
		case onePair:
			returnHand = threeOfAKind
		case twoPair:
			returnHand = fourOfAKind
		case fullHouse:
			returnHand = fiveOfAKind
		default:
			panic(fmt.Sprintf("Unexpected hand type with %d wildcards: %d", numberWildcards, currentHand))
		}

		return returnHand
	}

	if numberWildcards == 3 {
		switch currentHand {
		case threeOfAKind:
			returnHand = fourOfAKind
		case fullHouse:
			returnHand = fiveOfAKind
		default:
			panic(fmt.Sprintf("Unexpected hand type with %d wildcards: %d", numberWildcards, currentHand))
		}

		return returnHand
	}

	if numberWildcards == 4 {
		switch currentHand {
		case fourOfAKind:
			returnHand = fiveOfAKind
		default:
			panic(fmt.Sprintf("Unexpected hand type with %d wildcards: %d", numberWildcards, currentHand))
		}

		return returnHand
	}

	panic(fmt.Sprintf("Unexpected number of wildcards: %d", numberWildcards))
}

// constructHand assembles a pokerHand based on an input string. For example,
// "AAAQ4" would return a pokerHand object that identifies the string as a
// three-of-a-kind.
func constructHand(rawHand string) pokerHand {
	if len(rawHand) != 5 {
		panic("Improper input: " + rawHand)
	}

	retval := new(pokerHand)
	retval.raw = rawHand

	mapVals := make(map[byte]int)
	mapVals[rawHand[0]] = 1
	for i := 1; i < len(rawHand); i++ {
		char := rawHand[i]
		mapVals[char]++
	}

	// A map of length 1 indicates that all values are identical.
	if len(mapVals) == 1 {
		retval.value = fiveOfAKind
		return *retval
	}

	incrementValue := 0
	incrementValue = mapVals['J']

	// Compute each value as if J corresponds to a Jack instead of a wildcard
	// Joker. Then, pass the expected value into the computeWildcard method to
	// obtain the true value.
	if len(mapVals) == 5 {
		retval.value = computeWildcard(highCard, incrementValue)
		return *retval
	}
	if containsValue(mapVals, 4) {
		retval.value = computeWildcard(fourOfAKind, incrementValue)
		return *retval
	}
	if containsValue(mapVals, 3) {
		baseHandType := threeOfAKind
		if len(mapVals) == 2 {
			baseHandType = fullHouse
		}

		retval.value = computeWildcard(baseHandType, incrementValue)
		return *retval
	}
	if len(mapVals) == 3 {
		retval.value = computeWildcard(twoPair, incrementValue)
		return *retval
	}

	retval.value = computeWildcard(onePair, incrementValue)
	return *retval
}

func convertCardToInt(char byte) int {
	var num int
	if char >= '2' && char <= '9' {
		num = int(char - '0')
	} else if char == 'T' {
		num = 10
	} else if char == 'J' {
		num = 1
	} else if char == 'Q' {
		num = 12
	} else if char == 'K' {
		num = 13
	} else if char == 'A' {
		num = 14
	} else {
		panic("Invalid character found: " + string(char))
	}

	return num
}

// compareHands evaluates two poker hands and returns true if the first poker
// hand is losing against the second poker hand.
func compareHands(firstHand pokerHand, secondHand pokerHand) bool {
	if firstHand.value != secondHand.value {
		return firstHand.value < secondHand.value
	}

	for i := range firstHand.raw {
		val1 := convertCardToInt(firstHand.raw[i])
		val2 := convertCardToInt(secondHand.raw[i])

		if val1 != val2 {
			return val1 < val2
		}
	}

	panic("Two identical hands were found.")
}

func main() {
	fileLines := fileutils.LoadFile("input.txt")

	// Assemble the poker hands as a slice of pokerHand objects.
	handList := make([]pokerHand, 0)
	for _, line := range fileLines {
		values := strings.Fields(line)
		parsedHand := constructHand(values[0])
		parsedHand.bid, _ = strconv.Atoi(values[1])

		handList = append(handList, parsedHand)
	}

	// Sort the slice with the worst poker hand listed first.
	sort.Slice(handList, func(i, j int) bool {
		return compareHands(handList[i], handList[j])
	})

	totalWinnings := 0
	for i, handEntry := range handList {
		totalWinnings += (i + 1) * handEntry.bid
	}

	fmt.Printf("The total winnings across all the poker hands are %d.\n", totalWinnings)
}
