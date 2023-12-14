package day13

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

var trimRe = regexp.MustCompile(`(^[^\.#]+)|([^\.#]+$)`)

func Run() {
	fmt.Println("day 13")

	rawData := utils.ReadFile("13/data.txt")

	patterns := strings.Split(rawData, "\n\n")

	total1 := 0
	total2 := 0
	for _, rawPattern := range patterns {
		pattern, patternT, patternSlices := parseRawPattern(rawPattern)

		tempMR := findMirror(pattern, 100)

		mirrorRow := tempMR[0]

		withSmudge := mirrorRow
		if mirrorRow == 0 {
			tempMC := findMirror(patternT, 1)
			mirrorRowT := tempMC[0]
			withSmudge = mirrorRowT
		}

		total1 += withSmudge

		count := len(pattern) * len(patternT)
		noSmudgeSummary := 0
		for i := 0; i < count; i++ {
			noSmudge, noSmudgeT := fixSmudge(patternSlices, i)

			mirrorRow := 0
			temp1 := findMirror(noSmudge, 100)

			for _, mR := range temp1 {
				if mR != withSmudge {
					mirrorRow = mR
				}
			}

			noSmudgeSummary = mirrorRow
			if mirrorRow == 0 || noSmudgeSummary == withSmudge {
				mirrorRowT := 0
				temp2 := findMirror(noSmudgeT, 1)
				for _, mR := range temp2 {
					if mR != withSmudge {
						mirrorRowT = mR
					}
				}

				noSmudgeSummary = mirrorRowT
			}

			if noSmudgeSummary > 0 && noSmudgeSummary != withSmudge {
				total2 += noSmudgeSummary
				break
			}
		}

		if noSmudgeSummary == 0 {
			for i := 0; i < len(pattern); i++ {
				fmt.Println(pattern[i], i+1)
			}
			fmt.Println(withSmudge)
			fmt.Println()
		}
	}

	fmt.Println("total1:", total1)
	fmt.Println("total2:", total2)
}

func parseRawPattern(input string) ([]string, []string, [][]string) {
	temp := trimRe.ReplaceAllString(input, "")

	tempSplit := strings.Split(temp, "\n")

	tempPattern := [][]string{}
	pattern := []string{}
	for _, row := range tempSplit {
		pattern = append(pattern, row)
		cols := strings.Split(row, "")
		tempPattern = append(tempPattern, cols)
	}

	transposed := utils.TransposeSliceSlice(tempPattern)

	patternT := []string{}
	for _, rowS := range transposed {
		row := strings.Join(rowS, "")
		patternT = append(patternT, row)
	}

	return pattern, patternT, tempPattern
}

func findMirror(pattern []string, factor int) []int {
	allMirrors := []int{}
	for i := 1; i < len(pattern); i++ {
		isMirror := true
		width1 := i
		width2 := len(pattern) - i
		width := int(math.Min(float64(width1), float64(width2)))

		for j := 0; j < width; j++ {
			if pattern[i-j-1] != pattern[i+j] {
				isMirror = false
				break
			}
		}

		if isMirror {
			allMirrors = append(allMirrors, i*factor)
		}
	}

	if len(allMirrors) == 0 {
		return []int{0}
	}

	return allMirrors
}

func fixSmudge(slicePattern [][]string, smudge int) ([]string, []string) {
	temp := len(slicePattern[0])
	smudgeJ := smudge % temp

	// temp2 := len(slicePattern)
	smudgeI := int(math.Floor(float64(smudge) / float64(temp)))

	tempOutput := [][]string{}
	length := len(slicePattern[0])
	for i := 0; i < len(slicePattern); i++ {
		oo := make([]string, length)
		tempOutput = append(tempOutput, oo)
	}

	tempOutputT := [][]string{}
	length = len(slicePattern)
	for i := 0; i < len(slicePattern[0]); i++ {
		oo := make([]string, length)
		tempOutputT = append(tempOutputT, oo)
	}

	for i, inputI := range slicePattern {
		for j, inputJ := range inputI {
			temp := inputJ
			if i == smudgeI && j == smudgeJ {
				if temp == "#" {
					temp = "."
				} else {
					temp = "#"
				}
			}

			tempOutput[i][j] = temp
			tempOutputT[j][i] = temp
		}
	}

	output := []string{}
	for _, t := range tempOutput {
		temp := strings.Join(t, "")
		output = append(output, temp)
	}

	outputT := []string{}
	for _, t := range tempOutputT {
		temp := strings.Join(t, "")
		outputT = append(outputT, temp)
	}

	return output, outputT
}
