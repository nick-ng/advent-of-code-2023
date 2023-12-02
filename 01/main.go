package day01

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type token struct {
	M string
	V int
}

var numbers = []token{
	{
		M: "1", V: 1,
	},
	{
		M: "2", V: 2,
	},
	{
		M: "3", V: 3,
	},
	{
		M: "4", V: 4,
	},
	{
		M: "5", V: 5,
	},
	{
		M: "6", V: 6,
	},
	{
		M: "7", V: 7,
	},
	{
		M: "8", V: 8,
	},
	{
		M: "9", V: 9,
	},
	{
		M: "one", V: 1,
	},
	{
		M: "two", V: 2,
	},
	{
		M: "three", V: 3,
	},
	{
		M: "four", V: 4,
	},
	{
		M: "five", V: 5,
	},
	{
		M: "six", V: 6,
	},
	{
		M: "seven", V: 7,
	},
	{
		M: "eight", V: 8,
	},
	{
		M: "nine", V: 9,
	},
}

func Run() {
	fmt.Println("Day 1, Question 1")

	rawData := utils.ReadFile("01/data.txt")

	lines := strings.Split(rawData, "\n")

	total1 := 0
	total2 := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		// part 1
		re1a := regexp.MustCompile(`^[^\d]*(\d)`)

		firstCharacter := re1a.FindStringSubmatch(line)

		firstDigit, err := strconv.Atoi(firstCharacter[1])

		if err != nil {
			fmt.Println("couldn't match first character")
		}

		re1b := regexp.MustCompile(`(\d)[^\d]*$`)

		lastCharacter := re1b.FindStringSubmatch(line)

		lastDigit, err := strconv.Atoi(lastCharacter[1])

		if err != nil {
			fmt.Println("couldn't match last character")
		}

		number1 := firstDigit*10 + lastDigit

		total1 += number1

		// part2
		firstDigit2 := 0
		lastDigit2 := 0
		for i := range line {
			for _, t := range numbers {
				if len(line[i:]) < len(t.M) {
					continue
				}

				subString := line[i:(i + len(t.M))]

				if firstDigit2 == 0 && subString == t.M {
					firstDigit2 = t.V
				}

				if subString == t.M {
					lastDigit2 = t.V
				}
			}
		}

		temp := firstDigit2*10 + lastDigit2

		total2 += temp
	}

	fmt.Println("total1", total1)
	fmt.Println("total2", total2)
}
