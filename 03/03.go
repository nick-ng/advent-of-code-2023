package day03

import (
	"advent-of-code-2023/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CoOrd struct {
	Character string
	Row       int
	Col       int
}

type PartNumber struct {
	Id    int
	Value int
	Start CoOrd
	End   CoOrd
}

var symbolRe = regexp.MustCompile(`[^0-9\.]`)
var digitRe = regexp.MustCompile(`[0-9]`)

func Run() {
	fmt.Println("day 3")

	rawData := utils.ReadFile("03/data.txt")

	symbolCoOrds := []CoOrd{}
	numberCoOrds := []PartNumber{}
	startCoOrds := CoOrd{}

	rows := strings.Split(rawData, "\n")
	currentDigits := ""

	counter := 0
	for r, row := range rows {
		for c, col := range row {
			// 10: find all symbols
			character := string(col)
			symbolMatch := symbolRe.MatchString(character)

			if symbolMatch {
				symbolCoOrds = append(symbolCoOrds, CoOrd{
					Character: character,
					Row:       r,
					Col:       c,
				})
			}

			// 20: find all number
			digitMatch := digitRe.MatchString(character)

			if digitMatch {
				if len(currentDigits) == 0 {
					startCoOrds = CoOrd{
						Row: r,
						Col: c,
					}
				}

				currentDigits = fmt.Sprintf("%s%s", currentDigits, character)
			}

			if !digitMatch || c == len(row)-1 {
				if len(currentDigits) > 0 {
					partNumberValue, err := strconv.Atoi(currentDigits)

					if err != nil {
						fmt.Println("couldn't convert digits to a number", currentDigits)
						os.Exit(1)
					}

					numberCoOrds = append(numberCoOrds, PartNumber{
						Id:    counter,
						Value: partNumberValue,
						Start: startCoOrds,
						End: CoOrd{
							Row: r,
							Col: c - 1,
						},
					})

					counter++
					currentDigits = ""
				}
			}
		}
	}

	adjacentParts := []PartNumber{}
	gears := map[string][]int{}
	for _, symbolCoOrd := range symbolCoOrds {
		for _, partNumber := range numberCoOrds {
			rowDif := symbolCoOrd.Row - partNumber.Start.Row
			if rowDif >= -1 && rowDif <= 1 {
				if (symbolCoOrd.Col-1) <= partNumber.End.Col && (symbolCoOrd.Col+1) >= partNumber.Start.Col {
					newPart := true
					for _, pn := range adjacentParts {
						if pn.Id == partNumber.Id {
							newPart = false
							break
						}
					}

					if newPart {
						adjacentParts = append(adjacentParts, partNumber)

						// part 2
						if symbolCoOrd.Character == "*" {
							gearId := fmt.Sprintf("%d_%d", symbolCoOrd.Col, symbolCoOrd.Row)

							gears[gearId] = append(gears[gearId], partNumber.Value)
						}
					}
				}
			}
		}
	}

	total1 := 0
	for _, pn := range adjacentParts {
		total1 += pn.Value
	}

	total2 := 0
	for gearId, gg := range gears {
		if len(gg) == 2 {
			fmt.Println(gearId, gg)

			total2 += gg[0] * gg[1]
		}
	}

	fmt.Println("total1:", total1)
	fmt.Println("total2:", total2)
}
