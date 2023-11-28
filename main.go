package main

import (
	day01 "advent-of-code-2023/01"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	fmt.Println(args)

	switch os.Args[1] {
	case "1":
		{
			day01.Run()
		}
	default:
		{
			fmt.Println("Please specify a day.")
		}
	}
}
