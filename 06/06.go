package day06

import (
	"advent-of-code-2023/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("day 6")

	rawData := utils.ReadFile("06/data.txt")

	temp := strings.Split(rawData, "\n")

	times := utils.SpacesRe.Split(temp[0], -1)[1:]
	distances := utils.SpacesRe.Split(temp[1], -1)[1:]
	time2a := strings.Join(times, "")
	time2, err := strconv.Atoi(time2a)
	if err != nil {
		fmt.Println("cannot convert time2 to int", time2, err)
		os.Exit(1)
	}

	distance2a := strings.Join(distances, "")
	distance2, err := strconv.Atoi(distance2a)
	if err != nil {
		fmt.Println("cannot convert distance2 to int", distance2a, err)
		os.Exit(1)
	}

	total1 := 1
	for i := range times {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			fmt.Println("cannot convert time to int", times[i], err)
			os.Exit(1)
		}

		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			fmt.Println("cannot convert distance to int", distances[i], err)
			os.Exit(1)
		}

		inset := 0
		for j := 1; j < time; j++ {
			boatSpeed := j
			remainingTime := time - j

			boatDistance := boatSpeed * remainingTime

			if boatDistance > distance {
				inset = j
				break
			}
		}

		ways := time - (inset * 2) + 1

		total1 = total1 * ways
	}

	inset := 0
	for j := 1; j < time2; j++ {
		boatSpeed := j
		remainingTime := time2 - j

		boatDistance := boatSpeed * remainingTime

		if boatDistance > distance2 {
			inset = j
			break
		}
	}

	total2 := time2 - (inset * 2) + 1

	fmt.Println("total1:", total1)

	fmt.Println("part 2 time:", time2)
	fmt.Println("part 2 distance:", distance2)
	fmt.Println("total2:", total2)
}
