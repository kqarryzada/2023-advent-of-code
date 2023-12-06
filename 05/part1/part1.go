package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

type mapEntry struct {
	dest     int64
	source   int64
	rangeVal int64
}

var seedToSoilMap []mapEntry = make([]mapEntry, 0)
var soilToFertilizerMap []mapEntry = make([]mapEntry, 0)
var fertilizerToWaterMap []mapEntry = make([]mapEntry, 0)
var waterToLightMap []mapEntry = make([]mapEntry, 0)
var lightToTemperatureMap []mapEntry = make([]mapEntry, 0)
var temperatureToHumidityMap []mapEntry = make([]mapEntry, 0)
var humidityToLocationMap []mapEntry = make([]mapEntry, 0)

func parseLine(line string) *mapEntry {
	values := strings.Split(line, " ")
	if len(values) != 3 {
		panic("Error during parsing. Line is: " + line)
	}
	destVal, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		panic("Error during parsing. Line is: " + line)
	}
	sourceVal, err := strconv.ParseInt(values[1], 10, 64)
	if err != nil {
		panic("Error during parsing. Line is: " + line)
	}
	totalRange, err := strconv.ParseInt(values[2], 10, 64)
	if err != nil {
		panic("Error during parsing. Line is: " + line)
	}

	return &mapEntry{
		dest:     destVal,
		source:   sourceVal,
		rangeVal: totalRange,
	}
}

func initializeMaps(fileLines []string) {
	// This parsing is pretty manual and gnarly, but the input file is
	// structured in a very particular manner, and is consistent between the
	// example and the input files.
	i := 3
	for ; i < len(fileLines); i++ {
		line := fileLines[i]
		if len(line) == 0 {
			i += 2
			break
		}

		entry := parseLine(line)
		seedToSoilMap = append(seedToSoilMap, *entry)
	}

	for ; i < len(fileLines); i++ {
		line := fileLines[i]
		if len(line) == 0 {
			i += 2
			break
		}

		entry := parseLine(line)
		soilToFertilizerMap = append(soilToFertilizerMap, *entry)
	}

	for ; i < len(fileLines); i++ {
		line := fileLines[i]
		if len(line) == 0 {
			i += 2
			break
		}

		entry := parseLine(line)
		fertilizerToWaterMap = append(fertilizerToWaterMap, *entry)
	}

	for ; i < len(fileLines); i++ {
		line := fileLines[i]
		if len(line) == 0 {
			i += 2
			break
		}

		entry := parseLine(line)
		waterToLightMap = append(waterToLightMap, *entry)
	}

	for ; i < len(fileLines); i++ {
		line := fileLines[i]
		if len(line) == 0 {
			i += 2
			break
		}

		entry := parseLine(line)
		lightToTemperatureMap = append(lightToTemperatureMap, *entry)
	}

	for ; i < len(fileLines); i++ {
		line := fileLines[i]
		if len(line) == 0 {
			i += 2
			break
		}

		entry := parseLine(line)
		temperatureToHumidityMap = append(temperatureToHumidityMap, *entry)
	}

	for ; i < len(fileLines); i++ {
		line := fileLines[i]
		if len(line) == 0 {
			i += 2
			break
		}

		entry := parseLine(line)
		humidityToLocationMap = append(humidityToLocationMap, *entry)
	}

	sort.Slice(seedToSoilMap, func(a, b int) bool {
		return seedToSoilMap[a].source < seedToSoilMap[b].source
	})
	sort.Slice(soilToFertilizerMap, func(a, b int) bool {
		return soilToFertilizerMap[a].source < soilToFertilizerMap[b].source
	})
	sort.Slice(fertilizerToWaterMap, func(a, b int) bool {
		return fertilizerToWaterMap[a].source < fertilizerToWaterMap[b].source
	})
	sort.Slice(waterToLightMap, func(a, b int) bool {
		return waterToLightMap[a].source < waterToLightMap[b].source
	})
	sort.Slice(lightToTemperatureMap, func(a, b int) bool {
		return lightToTemperatureMap[a].source < lightToTemperatureMap[b].source
	})
	sort.Slice(temperatureToHumidityMap, func(a, b int) bool {
		return temperatureToHumidityMap[a].source < temperatureToHumidityMap[b].source
	})
	sort.Slice(humidityToLocationMap, func(a, b int) bool {
		return humidityToLocationMap[a].source < humidityToLocationMap[b].source
	})
}

func calculateValue(mapSlice []mapEntry, inputValue int64) int64 {
	if inputValue < mapSlice[0].source {
		return inputValue
	}
	lastMapEntry := mapSlice[len(mapSlice)-1]
	if inputValue > (lastMapEntry.source + lastMapEntry.rangeVal - 1) {
		return inputValue
	}

	var minValue mapEntry
	for _, val := range mapSlice {
		if val.source > inputValue {
			break
		}

		minValue = val
	}

	if inputValue > (minValue.source + minValue.rangeVal - 1) {
		return inputValue
	}

	diff := inputValue - minValue.source
	return diff + minValue.dest
}

func calculateLocationValue(seedNumber int64) int64 {
	soilValue := calculateValue(seedToSoilMap, seedNumber)
	fertilizerValue := calculateValue(soilToFertilizerMap, soilValue)
	waterValue := calculateValue(fertilizerToWaterMap, fertilizerValue)
	lightValue := calculateValue(waterToLightMap, waterValue)
	tempValue := calculateValue(lightToTemperatureMap, lightValue)
	humidityValue := calculateValue(temperatureToHumidityMap, tempValue)
	locationValue := calculateValue(humidityToLocationMap, humidityValue)

	return locationValue
}

func main() {
	fileLines := fileutils.LoadFile("input.txt")
	initializeMaps(fileLines)

	var minValue int64 = math.MaxInt64
	seedList := strings.Split(fileLines[0], " ")
	for i := 1; i < len(seedList); i++ {
		seedNumber, _ := strconv.ParseInt(seedList[i], 10, 64)
		value := calculateLocationValue(seedNumber)
		minValue = min(minValue, value)
	}

	fmt.Printf("The smallest location value is %d.\n", minValue)
}
