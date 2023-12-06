package day05

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	Destination int
	Source      int
	Width       int
}

func Run() {
	fmt.Println("Day 5")

	rawData := utils.ReadFile("05/data.txt")

	temp := strings.Split(rawData, "\n\n")

	seeds := []int{}
	seedToSoil := []Mapping{}
	soilToFertilizer := []Mapping{}
	fertilizerToWater := []Mapping{}
	waterToLight := []Mapping{}
	lightToTemperature := []Mapping{}
	temperatureToHumidity := []Mapping{}
	humidityToLocation := []Mapping{}

	seedRanges := [][]int{}

	for _, t := range temp {
		if strings.HasPrefix(t, "seeds:") {
			temp2 := strings.Split(t, " ")

			seedRange := []int{}

			for _, seedNumber := range temp2 {
				if strings.HasPrefix(seedNumber, "seeds:") {
					continue
				}

				seedInt, err := strconv.Atoi(seedNumber)

				if err != nil {
					fmt.Println("cannot convert seedNumber to int", seedInt)
					os.Exit(1)
				}

				seeds = append(seeds, seedInt)

				seedRange = append(seedRange, seedInt)

				if len(seedRange) == 2 {
					seedRanges = append(seedRanges, seedRange)

					seedRange = []int{}
				}
			}
		}

		if strings.HasPrefix(t, "seed-to-soil map:") {
			tempR := strings.Split(t, "\n")

			for _, seedToSoilRange := range tempR {
				if strings.HasPrefix(seedToSoilRange, "seed-to-soil map:") {
					continue
				}

				seedToSoil = applyRangeToMap(seedToSoil, seedToSoilRange)
			}
		}

		if strings.HasPrefix(t, "soil-to-fertilizer map:") {
			tempR := strings.Split(t, "\n")

			for _, soilToFertilizerRange := range tempR {
				if strings.HasPrefix(soilToFertilizerRange, "soil-to-fertilizer map:") {
					continue
				}

				soilToFertilizer = applyRangeToMap(soilToFertilizer, soilToFertilizerRange)
			}
		}

		if strings.HasPrefix(t, "fertilizer-to-water map:") {
			tempR := strings.Split(t, "\n")

			for _, fertilizerToWaterRange := range tempR {
				if strings.HasPrefix(fertilizerToWaterRange, "fertilizer-to-water map:") {
					continue
				}

				fertilizerToWater = applyRangeToMap(fertilizerToWater, fertilizerToWaterRange)
			}
		}

		if strings.HasPrefix(t, "water-to-light map:") {
			tempR := strings.Split(t, "\n")

			for _, waterToLightRange := range tempR {
				if strings.HasPrefix(waterToLightRange, "water-to-light map:") {
					continue
				}

				waterToLight = applyRangeToMap(waterToLight, waterToLightRange)
			}
		}

		if strings.HasPrefix(t, "light-to-temperature map:") {
			tempR := strings.Split(t, "\n")

			for _, lightToTemperatureRange := range tempR {
				if strings.HasPrefix(lightToTemperatureRange, "light-to-temperature map:") {
					continue
				}

				lightToTemperature = applyRangeToMap(lightToTemperature, lightToTemperatureRange)
			}
		}

		if strings.HasPrefix(t, "temperature-to-humidity map:") {
			tempR := strings.Split(t, "\n")

			for _, temperatureToHumidityRange := range tempR {
				if strings.HasPrefix(temperatureToHumidityRange, "temperature-to-humidity map:") {
					continue
				}

				temperatureToHumidity = applyRangeToMap(temperatureToHumidity, temperatureToHumidityRange)
			}
		}

		if strings.HasPrefix(t, "humidity-to-location map:") {
			tempR := strings.Split(t, "\n")

			for _, humidityToLocationRange := range tempR {
				if strings.HasPrefix(humidityToLocationRange, "humidity-to-location map:") {
					continue
				}

				humidityToLocation = applyRangeToMap(humidityToLocation, humidityToLocationRange)
			}
		}
	}

	minLocation := math.MaxInt

	for _, seed := range seeds {
		soil := getDestination(seedToSoil, seed)
		fertilizer := getDestination(soilToFertilizer, soil)
		water := getDestination(fertilizerToWater, fertilizer)
		light := getDestination(waterToLight, water)
		temperature := getDestination(lightToTemperature, light)
		humidity := getDestination(temperatureToHumidity, temperature)
		location := getDestination(humidityToLocation, humidity)

		if location < minLocation {
			minLocation = location
		}
	}

	fmt.Println("min location 1:", minLocation)

	soilRanges := getDestinationFromRange(seedToSoil, seedRanges)
	fertilizerRanges := getDestinationFromRange(soilToFertilizer, soilRanges)
	waterRanges := getDestinationFromRange(fertilizerToWater, fertilizerRanges)
	lightRanges := getDestinationFromRange(waterToLight, waterRanges)
	temperatureRanges := getDestinationFromRange(lightToTemperature, lightRanges)
	humidityRanges := getDestinationFromRange(temperatureToHumidity, temperatureRanges)
	locationRanges := getDestinationFromRange(humidityToLocation, humidityRanges)

	minLocation2 := math.MaxInt

	for _, locationRange := range locationRanges {
		if locationRange[0] < minLocation2 && locationRange[1] > 0 {
			minLocation2 = locationRange[0]
		}
	}

	fmt.Println("min location 2:", minLocation2)
}

func applyRangeToMap(mappings []Mapping, rangeString string) []Mapping {
	if len(rangeString) == 0 {
		return mappings
	}

	mapStrings := strings.Split(rangeString, " ")
	mapInts := []int{}

	for _, seedToSoilString := range mapStrings {
		t, err := strconv.Atoi(seedToSoilString)

		if err != nil {
			fmt.Println("couldn't convert seedToSoilString to int", seedToSoilString)
			os.Exit(1)
		}

		mapInts = append(mapInts, t)
	}

	destination := mapInts[0]
	source := mapInts[1]
	width := mapInts[2]

	mappings = append(mappings, Mapping{
		Destination: destination,
		Source:      source,
		Width:       width,
	})

	return mappings
}

func getDestination(mappings []Mapping, source int) int {
	for _, mapping := range mappings {
		if mapping.Source <= source && mapping.Source+mapping.Width >= source {
			offset := source - mapping.Source

			output := mapping.Destination + offset

			return output
		}
	}

	return source
}

func getDestinationFromRange(mappings []Mapping, inputRanges [][]int) [][]int {
	outputRanges := [][]int{}

	for _, inputRange := range inputRanges {
		inputSource := inputRange[0]
		inputWidth := inputRange[1]

		for inputWidth > 0 {
			nextRangeSource := math.MaxInt
			for _, mapping := range mappings {
				if mapping.Source <= inputSource && mapping.Source+mapping.Width >= inputSource {

					offset := inputSource - mapping.Source
					maxWidth := mapping.Width - offset
					start := mapping.Destination + offset

					if inputWidth > maxWidth {
						inputWidth = inputWidth - maxWidth
						inputSource = inputSource + maxWidth
						outRange := []int{start, maxWidth}
						outputRanges = append(outputRanges, outRange)
					} else {
						outRange := []int{start, inputWidth}
						outputRanges = append(outputRanges, outRange)
						inputSource = inputSource + inputWidth // for consistency. the loop will exit
						inputWidth = 0
					}

					continue
				}

				if mapping.Source > inputSource && nextRangeSource > mapping.Source {
					nextRangeSource = mapping.Source
				}
			}

			// It wasn't in one of the ranges so each number maps to the same number
			offset := 0
			maxWidth := nextRangeSource - inputSource
			start := inputSource + offset // actually just inputSource

			if inputWidth > maxWidth {
				inputWidth = inputWidth - maxWidth
				inputSource = inputSource + maxWidth
				outRange := []int{start, maxWidth}
				outputRanges = append(outputRanges, outRange)
			} else {
				outRange := []int{start, inputWidth}
				outputRanges = append(outputRanges, outRange)
				inputSource = inputSource + inputWidth // for consistency. the loop will exit
				inputWidth = 0
			}
		}
	}

	return outputRanges
}
