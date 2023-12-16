package day15

import (
	"advent-of-code-2023/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var stepRe = regexp.MustCompile(`^(.+)([=-])(\d*)$`)

func Run() {
	fmt.Println("day 15")

	rawData := utils.ReadFile("15/data.txt")

	trimmedData := strings.TrimSpace(rawData)
	steps := strings.Split(trimmedData, ",")

	total1 := 0
	boxes := make([][]string, 256)
	for _, step := range steps {
		if len(step) == 0 {
			continue
		}
		stepHash := stringToHash(step)
		total1 += stepHash

		match := stepRe.FindStringSubmatch(step)

		if len(match) < 4 {
			fmt.Println("step did not match enough", step)
			continue
		}

		lensName := match[1]
		operation := match[2]

		boxNumber := stringToHash(lensName)

		if operation == "=" {
			focalLength := match[3]
			lensId := fmt.Sprintf("%s %s", lensName, focalLength)
			if len(boxes[boxNumber]) == 0 {
				boxes[boxNumber] = append(boxes[boxNumber], lensId)
				continue
			}

			lensIndex := -1
			for i, lens := range boxes[boxNumber] {
				if strings.HasPrefix(lens, lensName) {
					lensIndex = i
				}
			}

			if lensIndex < 0 {
				boxes[boxNumber] = append(boxes[boxNumber], lensId)
			} else {
				boxes[boxNumber][lensIndex] = lensId
			}
		} else if operation == "-" {
			if len(boxes[boxNumber]) == 0 {
				continue
			}

			lensIndex := -1
			for i, lens := range boxes[boxNumber] {
				if strings.HasPrefix(lens, lensName) {
					lensIndex = i
				}
			}

			if lensIndex < 0 {
				continue
			}

			part1 := boxes[boxNumber][:lensIndex]
			part2 := boxes[boxNumber][lensIndex+1:]

			boxes[boxNumber] = append(part1, part2...)
		}
	}

	fmt.Println("total1:", total1)

	total2 := 0

	for i, lenses := range boxes {
		for j, lens := range lenses {
			temp := strings.Split(lens, " ")

			focalLength, err := strconv.Atoi(temp[1])

			if err != nil {
				fmt.Println("error converting focalLength to int", temp[1], err)
				os.Exit(1)
			}

			focusingPower := (i + 1) * (j + 1) * focalLength

			total2 += focusingPower
		}
	}

	fmt.Println("total2:", total2)
}

func stringToHash(myString string) int {
	currentValue := 0

	for _, r := range myString {
		currentValue += int(r)

		currentValue = currentValue * 17
		currentValue = currentValue % 256
	}

	return currentValue
}
