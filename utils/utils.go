package utils

import (
	"fmt"
	"os"
	"regexp"
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
