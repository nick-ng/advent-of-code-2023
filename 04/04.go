package day04

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

var split1Re = regexp.MustCompile(` *(:|\|) *`)
var spacesRe = regexp.MustCompile(` +`)

func Run() {
	fmt.Println("Day 4")

	rawData := utils.ReadFile("04/data.txt")

	cards := strings.Split(rawData, "\n")

	total1 := 0.0
	total2 := 0

	cardCopies := map[int]int{}

	for i, card := range cards {
		if len(card) == 0 {
			continue
		}

		temp1 := split1Re.Split(card, -1)

		if len(temp1) != 3 {
			fmt.Println(i, "incorrect number of items in card", temp1)

			os.Exit(1)
		}

		winningNumbers := spacesRe.Split(temp1[1], -1)
		myNumbers := spacesRe.Split(temp1[2], -1)

		matches := 0
		for _, winning := range winningNumbers {
			for _, my := range myNumbers {
				if winning == my {
					matches += 1
				}
			}
		}

		// part 1
		if matches > 0.0 {
			result := math.Pow(2, float64(matches-1))
			total1 += result
		}

		// part 2
		copies := 1 + cardCopies[i]
		total2 += copies

		for j := 0; j < matches; j++ {
			cardCopies[i+j+1] += copies
		}
	}

	fmt.Println("total1:", total1)
	fmt.Println("total2:", total2)
}
