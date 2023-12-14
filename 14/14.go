package day14

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"
)

type RockLogEntry struct {
	Rocks  string
	Cycles int
	Load   int
}

func Run() {
	fmt.Println("day 14")

	rawData := utils.ReadFile("14/data.txt")

	rows := strings.Split(rawData, "\n")

	columns := utils.TransposeSliceString(rows)

	total1 := 0
	for _, column := range columns {
		subsections := strings.Split(column, "#")
		newSubsections := []string{}
		for _, subsection := range subsections {
			oCount := 0
			for _, char := range subsection {
				if char == 'O' {
					oCount++
				}
			}

			newSubsection := utils.PadStringStart("", oCount, "O")
			newSubsection = utils.PadStringEnd(newSubsection, len(subsection), ".")

			newSubsections = append(newSubsections, newSubsection)
		}

		temp2 := strings.Join(newSubsections, "#")
		rolledColumn := strings.Split(temp2, "")

		columnValue := 0
		for i, char := range rolledColumn {
			if char == "O" {
				weight := len(rolledColumn) - i
				columnValue += weight
			}
		}

		total1 += columnValue
	}

	// part 2
	rows2 := []string{}
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		rows2 = append(rows2, row)
	}

	fullCycles := 1000000000
	rockLog := []RockLogEntry{}
	cycleCount1 := 0
	cycleCount2 := 0
	for i := 0; i < 1000; i++ {
		rows2 = spinCycle(rows2)

		rocks := strings.Join(rows2, "\n")

		isRepeat := false
		for _, rockLogEntry := range rockLog {
			if rockLogEntry.Rocks == rocks {
				isRepeat = true
				cycleCount1 = rockLogEntry.Cycles
				cycleCount2 = i + 1

				break
			}
		}

		if isRepeat {
			break
		}

		newLoad := calculateLoad(rows2)

		rockLogEntry := RockLogEntry{
			Rocks:  rocks,
			Cycles: i + 1,
			Load:   newLoad,
		}

		rockLog = append(rockLog, rockLogEntry)
	}

	fmt.Println("total1:", total1)

	fmt.Println("cycleCount1", cycleCount1)
	fmt.Println("cycleCount2", cycleCount2)
	newStart := cycleCount1
	fmt.Println("new start", newStart)
	newFullCycles := fullCycles - newStart
	fmt.Println("new full cycles", newFullCycles)

	cycleLength := cycleCount2 - cycleCount1

	fmt.Println("cycle length", cycleLength)

	rockCycle := rockLog[cycleCount1-1 : cycleCount2-1]

	modCycle := newFullCycles % cycleLength

	finalRockLogEntry := rockCycle[modCycle]

	finalLoad := finalRockLogEntry.Load
	fmt.Println("finalLoad", finalLoad)
	fmt.Println("--")
}

func rollToStart(rows []string) []string {
	rolledRows := []string{}
	for _, row := range rows {
		subsections := strings.Split(row, "#")
		newSubsections := []string{}
		for _, subsection := range subsections {
			oCount := 0
			for _, char := range subsection {
				if char == 'O' {
					oCount++
				}
			}

			newSubsection := utils.PadStringStart("", oCount, "O")
			newSubsection = utils.PadStringEnd(newSubsection, len(subsection), ".")

			newSubsections = append(newSubsections, newSubsection)
		}

		rolledRow := strings.Join(newSubsections, "#")

		rolledRows = append(rolledRows, rolledRow)
	}

	return rolledRows
}

func rollToEnd(rows []string) []string {
	rolledRows := []string{}
	for _, row := range rows {
		subsections := strings.Split(row, "#")
		newSubsections := []string{}
		for _, subsection := range subsections {
			oCount := 0
			for _, char := range subsection {
				if char == 'O' {
					oCount++
				}
			}

			newSubsection := utils.PadStringEnd("", oCount, "O")
			newSubsection = utils.PadStringStart(newSubsection, len(subsection), ".")

			newSubsections = append(newSubsections, newSubsection)
		}

		rolledRow := strings.Join(newSubsections, "#")

		rolledRows = append(rolledRows, rolledRow)
	}

	return rolledRows
}

func spinCycle(rows []string) []string {
	// north: transpose to columns, roll to start, transpose to rows
	temp := utils.TransposeSliceString(rows)
	temp = rollToStart(temp)
	temp = utils.TransposeSliceString(temp)

	// west: roll to start
	temp = rollToStart(temp)

	// south: transpose, roll to end, transpose
	temp = utils.TransposeSliceString(temp)
	temp = rollToEnd(temp)
	temp = utils.TransposeSliceString(temp)

	// east: roll to end
	temp = rollToEnd(temp)

	return temp
}

func calculateLoad(rows []string) int {
	columns := utils.TransposeSliceString(rows)

	totalLoad := 0
	for _, col := range columns {
		columnValue := 0
		for i, char := range col {
			if char == 'O' {
				weight := len(col) - i
				columnValue += weight
			}
		}

		totalLoad += columnValue
	}

	return totalLoad
}
