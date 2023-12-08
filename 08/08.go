package day08

import (
	"advent-of-code-2023/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type nodeMapType struct {
	Node  string
	Left  string
	Right string
}

var nodeRe = regexp.MustCompile(`([0-9A-Z]+) = \(([0-9A-Z]+), ([0-9A-Z]+)\)`)

func Run() {
	fmt.Println("day 8")

	primesData := utils.ReadFile("08/primes.txt")

	primesData = strings.ReplaceAll(primesData, "\n", " ")

	tempPrimes := utils.SpacesRe.Split(primesData, -1)

	primes := []int{}

	for _, rawPrime := range tempPrimes {
		if len(rawPrime) == 0 {
			continue
		}
		prime, err := strconv.Atoi(rawPrime)

		if err != nil {
			fmt.Println("cannot convert prime to int", rawPrime, err)
			os.Exit(1)
		}

		primes = append(primes, prime)
	}

	rawData := utils.ReadFile("08/data.txt")

	temp1 := strings.Split(rawData, "\n\n")

	directions := strings.Split(temp1[0], "")
	nodes := strings.Split(temp1[1], "\n")

	nodeMap := map[string]nodeMapType{}
	currentNodes := []string{}
	for _, node := range nodes {
		if len(node) == 0 {
			continue
		}

		temp := nodeRe.FindAllStringSubmatch(node, -1)

		if len(temp[0]) < 4 {
			fmt.Println("cannot parse node", node)
			continue
		}

		n := temp[0][1]
		l := temp[0][2]
		r := temp[0][3]

		nodeMap[n] = nodeMapType{
			Node:  n,
			Left:  l,
			Right: r,
		}

		if strings.HasSuffix(n, "A") {
			currentNodes = append(currentNodes, n)
		}
	}

	fmt.Println("startingNodes", currentNodes)

	periods := map[int][]int{}
	keepGoing := true
	steps1 := 0
	for keepGoing {
		for _, direction := range directions {
			if direction == "L" {
				for i, currentNode := range currentNodes {
					currentNodes[i] = nodeMap[currentNode].Left
				}
				steps1++
			} else if direction == "R" {
				for i, currentNode := range currentNodes {
					currentNodes[i] = nodeMap[currentNode].Right
				}
				steps1++
			}

			keepGoing = false
			for i, currentNode := range currentNodes {
				if !strings.HasSuffix(currentNode, "Z") {
					keepGoing = true
				} else {
					fmt.Println(steps1, i, currentNodes, currentNode)
					periods[i] = append(periods[i], steps1)
				}
			}

			if len(periods) == len(currentNodes) {
				keepGoing = false
				break
			}

			if !keepGoing {
				break
			}

			if steps1%1000001 == 0 {
				fmt.Println("steps:", steps1)
				fmt.Println(currentNodes)
			}
		}
	}

	fmt.Println("final steps1:", steps1)

	// find least common multiple with table-method
	periods2 := []int{}
	for _, p := range periods {
		periods2 = append(periods2, p[0])
	}

	fmt.Println("periods before", periods2)
	keepGoing = true

	table := []int{}
	for keepGoing {
		for _, prime := range primes {
			shouldBreak := false
			keepGoing = false

			divided := false

			for i, p := range periods2 {
				if p == 1 {
					// do nothing
				} else if p%prime == 0 {
					divided = true
					newP := p / prime
					periods2[i] = newP
					shouldBreak = true
					if newP != 1 {
						shouldBreak = true
					}
					if p != 1 {
						keepGoing = true
					}
				} else {
					periods2[i] = p

					keepGoing = true
				}
			}

			if divided {
				table = append(table, prime)
			}

			if shouldBreak {
				break
			}
		}
	}

	fmt.Println("table:", table)
	fmt.Println("periods after", periods2)

	product := 1
	for _, t := range table {
		product *= t
	}

	fmt.Println("product:", product)
}
