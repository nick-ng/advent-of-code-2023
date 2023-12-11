package day11

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"strings"
)

func Run() {
	fmt.Println("day 11")

	rawData := utils.ReadFile("11/data.txt")

	rows := strings.Split(rawData, "\n")

	universe := [][]string{}
	emptyRows := []int{}
	for i, tempRow := range rows {
		if len(tempRow) == 0 {
			continue
		}

		row := strings.Split(tempRow, "")

		universe = append(universe, row)

		if !utils.SliceContainsString(row, "#") {
			emptyRows = append(emptyRows, i)
		}
	}

	// find empty columns
	emptyColumns := []int{}
	for i := 0; i < len(universe[0]); i++ {
		isEmpty := true

		for _, row := range universe {
			if row[i] == "#" {
				isEmpty = false
			}
		}

		if isEmpty {
			emptyColumns = append(emptyColumns, i)
		}
	}

	emptyRow := []string{}

	for i := 0; i < len(universe[0])+len(emptyColumns); i++ {
		emptyRow = append(emptyRow, ".")
	}

	// expand empty rows
	expandedUniverse1 := [][]string{}
	expandedRows := 0
	galaxies2 := [][]int{}
	for i, row := range universe {
		newRow := []string{}
		for j, col := range row {
			newRow = append(newRow, col)

			if utils.SliceContainsInt(emptyColumns, j) {
				newRow = append(newRow, ".")
			}

			if col == "#" {
				galaxies2 = append(galaxies2, []int{i, j})
			}
		}

		expandedUniverse1 = append(expandedUniverse1, newRow)
		if utils.SliceContainsInt(emptyRows, i) {
			expandedRows++

			newEmptyRow := make([]string, len(emptyRow))
			copy(newEmptyRow, emptyRow)
			expandedUniverse1 = append(expandedUniverse1, newEmptyRow)
		}
	}

	galaxies := [][]int{}
	for i, row := range expandedUniverse1 {
		for j, col := range row {
			if col == "#" {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	total1 := 0
	for i, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies[i+1:] {
			differenceX := galaxy1[0] - galaxy2[0]
			horizontal := math.Abs(float64(differenceX))

			differenceY := galaxy1[1] - galaxy2[1]
			vertical := math.Abs(float64(differenceY))

			manhattan := int(horizontal) + int(vertical)

			total1 += manhattan
		}
	}

	emptyAmount := 1000000
	total2 := 0
	for i, galaxy1 := range galaxies2 {
		for _, galaxy2 := range galaxies2[i+1:] {
			start1 := int(math.Min(float64(galaxy1[0]), float64(galaxy2[0])))
			end1 := int(math.Max(float64(galaxy1[0]), float64(galaxy2[0])))

			diff1 := 0
			for ii := start1; ii < end1; ii++ {
				if utils.SliceContainsInt(emptyRows, ii) {
					diff1 += emptyAmount
				} else {
					diff1++
				}
			}

			start2 := int(math.Min(float64(galaxy1[1]), float64(galaxy2[1])))
			end2 := int(math.Max(float64(galaxy1[1]), float64(galaxy2[1])))

			diff2 := 0
			for jj := start2; jj < end2; jj++ {
				if utils.SliceContainsInt(emptyColumns, jj) {
					diff2 += emptyAmount
				} else {
					diff2++
				}
			}

			manhattan2 := diff1 + diff2

			total2 += manhattan2
		}
	}

	fmt.Println("total1:", total1)
	fmt.Println("total2:", total2)
}
