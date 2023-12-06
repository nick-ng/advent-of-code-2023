package day06

import (
	"advent-of-code-2023/utils"
	"fmt"
)

type Mapping struct {
	Destination int
	Source      int
	Width       int
}

func Run() {
	fmt.Println("Day 6")

	rawData := utils.ReadFile("06/data.txt")

	fmt.Println(rawData)
}
