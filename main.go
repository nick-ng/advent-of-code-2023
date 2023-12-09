package main

import (
	day01 "advent-of-code-2023/01"
	day02 "advent-of-code-2023/02"
	day03 "advent-of-code-2023/03"
	day04 "advent-of-code-2023/04"
	day05 "advent-of-code-2023/05"
	day06 "advent-of-code-2023/06"
	day07 "advent-of-code-2023/07"
	day08 "advent-of-code-2023/08"
	day09 "advent-of-code-2023/09"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

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
	case "3":
		{
			day03.Run()
		}
	case "4":
		{
			day04.Run()
		}
	case "5":
		{
			day05.Run()
		}
	case "6":
		{
			day06.Run()
		}
	case "7":
		{
			day07.Run()
			day07.RunJokers()
		}
	case "8":
		{
			day08.Run()
		}
	case "9":
		{
			day09.Run()
		}
	default:
		{
			fmt.Println("Please specify a day.")
			os.Exit(1)
		}
	}

	finish := time.Now()
	duration := finish.Sub(start)

	fmt.Println("duration:", duration)
}
