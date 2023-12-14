package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var SpacesRe = regexp.MustCompile(` +`)

func ReadFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("cannot read file:", path)
		fmt.Println(err)
		os.Exit(1)
	}

	return string(bytes)
}

func SliceContainsString(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}

func SliceContainsInt(haystack []int, needle int) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}

func PadStringStart(s string, minLength int, padding string) string {
	newString := s

	for len(newString) < minLength {
		newString = padding + newString
	}

	return newString
}

func PadStringEnd(s string, minLength int, padding string) string {
	newString := s

	for len(newString) < minLength {
		newString = newString + padding
	}

	return newString
}

func TransposeSliceSlice(input [][]string) [][]string {
	output := make([][]string, len(input[0]))
	for _, row := range input {
		for i, col := range row {
			output[i] = append(output[i], col)
		}
	}

	return output
}

func TransposeSliceString(input []string) []string {
	columns := make([][]string, len(input[0]))
	for _, row := range input {
		cols := strings.Split(row, "")
		for i, col := range cols {
			columns[i] = append(columns[i], col)
		}
	}

	output := []string{}
	for _, col := range columns {
		newColumn := strings.Join(col, "")
		output = append(output, newColumn)
	}

	return output
}
