package utils

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Couldn't read file:", path)
		fmt.Println(err)
		os.Exit(1)
	}

	return string(bytes)
}
