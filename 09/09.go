package day09

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("day 9")

	rawData := utils.ReadFile("09/data.txt")

	rawSequences := strings.Split(rawData, "\n")

	total1 := 0
	total2 := 0
	for _, rawSequence := range rawSequences {
		if len(rawSequence) == 0 {
			continue
		}

		temp := utils.SpacesRe.Split(rawSequence, -1)
		sequence := []int{}
		for _, rawValue := range temp {
			value, err := strconv.Atoi(rawValue)

			if err != nil {
				fmt.Println("cannot convert rawValue to int", rawValue, err)
				os.Exit(1)
			}

			sequence = append(sequence, value)
		}

		prevValue, nextValue := extrapolate(sequence)

		total1 += nextValue
		total2 += prevValue
	}

	// fmt.Println(sequences)
	fmt.Println("total1:", total1)
	fmt.Println("total2:", total2)
}

// Extrapolates a sequence
func extrapolate(sequence []int) (int, int) {
	differenceSequence := []int{}

	if len(sequence) < 2 {
		fmt.Println("sequence too short", sequence)
		os.Exit(10)
	}

	maxDifference := 0
	minDifference := math.MaxInt
	for i := 1; i < len(sequence); i++ {
		difference := sequence[i] - sequence[i-1]

		differenceSequence = append(differenceSequence, difference)

		if difference > maxDifference {
			maxDifference = difference
		}
		if difference < minDifference {
			minDifference = difference
		}
	}

	lastValue := sequence[len(sequence)-1]

	if maxDifference == 0 && minDifference == 0 {
		nextValue := lastValue
		return nextValue, nextValue
	}

	prevDifference, nextDifference := extrapolate(differenceSequence)
	nextValue := lastValue + nextDifference
	prevValue := sequence[0] - prevDifference
	return prevValue, nextValue
}
