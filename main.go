package main

import (
	day01 "advent-of-code-2023/01"
	day02 "advent-of-code-2023/02"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a day.")
		os.Exit(1)
	}

	args := os.Args[1:]

	switch args[0] {
	case "1":
		{
			day01.Run()
		}
	case "2":
		{
			day02.Run()
		}
	default:
		{
			fmt.Println("Please specify a day.")
			os.Exit(1)
		}
	}
}
