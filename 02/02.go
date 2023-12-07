package day02

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var maxRed1 = 12
var maxGreen1 = 13
var maxBlue1 = 14

func Run() {
	fmt.Println("day 2")

	rawData := utils.ReadFile("02/data.txt")

	lines := strings.Split(rawData, "\n")

	// part 1 & 2
	total1 := 0
	total2 := 0

	gameIdRe := regexp.MustCompile(`Game (\d+)`)
	redRe := regexp.MustCompile(`(\d+) red`)
	greenRe := regexp.MustCompile(`(\d+) green`)
	blueRe := regexp.MustCompile(`(\d+) blue`)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		temp1 := strings.Split(line, ":")

		if len(temp1) < 2 {
			fmt.Println("invalid game string:", line)
			continue
		}

		gameIdMatch := gameIdRe.FindStringSubmatch(line)

		if len(gameIdMatch) < 2 {
			fmt.Println("not enough elements:", line)
			continue
		}
		gameId, err := strconv.ParseInt(gameIdMatch[1], 10, 0)

		if err != nil {
			fmt.Println("cannot parse int:", err)
			continue
		}

		rounds := strings.Split(temp1[1], ";")

		can := true

		redMin := 0
		greenMin := 0
		blueMin := 0

		for _, round := range rounds {
			redMatch := redRe.FindStringSubmatch(round)

			if len(redMatch) >= 2 {
				redCount, err := strconv.Atoi(redMatch[1])

				if err != nil {
					continue
				}

				if redCount > maxRed1 {
					can = false
				}

				if redCount > redMin {
					redMin = redCount
				}
			}

			greenMatch := greenRe.FindStringSubmatch(round)

			if len(greenMatch) >= 2 {
				greenCount, err := strconv.Atoi(greenMatch[1])

				if err != nil {
					continue
				}

				if greenCount > maxGreen1 {
					can = false
				}

				if greenCount > greenMin {
					greenMin = greenCount
				}
			}

			blueMatch := blueRe.FindStringSubmatch(round)

			if len(blueMatch) >= 2 {
				blueCount, err := strconv.Atoi(blueMatch[1])

				if err != nil {
					continue
				}

				if blueCount > maxBlue1 {
					can = false
				}

				if blueCount > blueMin {
					blueMin = blueCount
				}
			}
		}

		power := redMin * greenMin * blueMin

		total2 += power

		if can {
			total1 += int(gameId)
		}
	}

	fmt.Println("total1", total1)
	fmt.Println("total2", total2)
}
