package day12

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("day 12")

	rawData := utils.ReadFile("12/data.txt")

	rows := strings.Split(rawData, "\n")

	total1 := 0
	total2 := 0
	for i, row := range rows {
		if len(row) == 0 {
			continue
		}

		temp := strings.Split(row, " ")
		springConditions := temp[0]
		springConditions2 := temp[0] + "?" + temp[0] + "?" + temp[0] + "?" + temp[0] + "?" + temp[0]

		temp2 := strings.Split(temp[1], ",")

		springConditionsGroups := []int{}

		totalDamaged := 0
		for _, t := range temp2 {
			value, err := strconv.Atoi(t)

			if err != nil {
				fmt.Println("cannot convert to int", t, err)
				os.Exit(1)
			}

			springConditionsGroups = append(springConditionsGroups, value)
			totalDamaged += value
		}

		springConditionsGroups2 := append(springConditionsGroups, springConditionsGroups...)
		springConditionsGroups2 = append(springConditionsGroups2, springConditionsGroups...)
		springConditionsGroups2 = append(springConditionsGroups2, springConditionsGroups...)
		springConditionsGroups2 = append(springConditionsGroups2, springConditionsGroups...)

		combinations := getCombinations(springConditions, totalDamaged)
		combinations2 := getCombinations(springConditions2, totalDamaged*5)

		validCombinations := 0
		for _, combination := range combinations {
			overall := validateRow2(combination, springConditionsGroups)

			if overall {
				validCombinations++
			}
		}

		// fmt.Println(i+1, springConditions, validCombinations, "/", len(combinations))
		total1 += validCombinations

		validCombinations2 := 0
		for _, combination := range combinations2 {
			overall := validateRow2(combination, springConditionsGroups2)

			if overall {
				validCombinations2++
			}
		}

		fmt.Println(i+1, validCombinations2, "/", len(combinations2))
		total2 += validCombinations2
	}

	fmt.Println("total1:", total1)
	fmt.Println("total2:", total2)
}

var dotsRe = regexp.MustCompile(`\.+`)
var dotsStartRe = regexp.MustCompile(`^\.+`)
var dotsEndRe = regexp.MustCompile(`\.+$`)

func validateRow(springConditions string, springConditionsGroups []int) ([]bool, bool) {
	result := []bool{}

	temp := dotsStartRe.ReplaceAllString(springConditions, "")
	temp = dotsEndRe.ReplaceAllString(temp, "")

	splitGroups := dotsRe.Split(temp, -1)

	overall := true
	for i := 0; i < len(springConditionsGroups); i++ {
		if len(splitGroups) <= i {
			result = append(result, false)
			overall = false
			continue
		}

		group := splitGroups[i]

		if len(group) != springConditionsGroups[i] {
			result = append(result, false)
			overall = false
			continue
		}

		result = append(result, true)
	}

	if len(splitGroups) != len(springConditionsGroups) {
		overall = false
	}

	return result, overall
}

func validateRow2(springConditions string, springConditionsGroups []int) bool {
	temp := dotsStartRe.ReplaceAllString(springConditions, "")
	temp = dotsEndRe.ReplaceAllString(temp, "")

	splitGroups := dotsRe.Split(temp, -1)

	if len(splitGroups) != len(springConditionsGroups) {
		return false
	}
	for i := 0; i < len(springConditionsGroups); i++ {
		if len(splitGroups) <= i {
			return false
		}

		group := splitGroups[i]

		if len(group) != springConditionsGroups[i] {
			return false
		}
	}

	return true
}

func getCombinations(springConditions string, totalDamaged int) []string {
	individualSprings := strings.Split(springConditions, "")
	questionMarkCount := 0
	knownDamaged := 0
	for _, c := range individualSprings {
		if c == "?" {
			questionMarkCount++
		} else if c == "#" {
			knownDamaged++
		}
	}

	maxValue := int(math.Pow(2, float64(questionMarkCount)))

	combinations := []string{}
	for i := 0; i < maxValue; i++ {
		binary := strconv.FormatInt(int64(i), 2)

		b2 := utils.PadStringStart(binary, questionMarkCount, "0")

		temp := strings.Split(b2, "")

		newCombination := ""
		damagedCount := 0
		for j := 0; j < len(individualSprings); j++ {
			char := individualSprings[j]
			if char == "?" {
				if temp[0] == "1" {
					char = "#"
				} else {
					char = "."
				}

				temp = temp[1:]
			}
			if char == "#" {
				damagedCount++
			}

			newCombination += char
		}

		if damagedCount == totalDamaged {
			combinations = append(combinations, newCombination)
		}
	}

	return combinations
}
