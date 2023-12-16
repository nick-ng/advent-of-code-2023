package day16

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"
)

type Direction struct {
	H int
	V int
	C string
}

type Position struct {
	X int
	Y int
}

type LightParticle struct {
	Direction Direction
	Position  Position
	Hash      string
}

var up = Direction{
	H: 0,
	V: -1,
	C: "^",
}

var down = Direction{
	H: 0,
	V: 1,
	C: "v",
}

var left = Direction{
	H: -1,
	V: 0,
	C: "<",
}

var right = Direction{
	H: 1,
	V: 0,
	C: ">",
}

func Run() {
	fmt.Println("day 16")

	rawData := utils.ReadFile("16/data.txt")

	rows := strings.Split(rawData, "\n")

	contraptionLayout := [][]string{}
	for _, rawRow := range rows {
		if len(rawRow) == 0 {
			continue
		}

		row := strings.Split(rawRow, "")

		contraptionLayout = append(contraptionLayout, row)
	}

	lights := []LightParticle{
		{
			Position: Position{
				X: 0,
				Y: 0,
			},
			Direction: right,
		},
	}

	total1 := calculateLitSquares(contraptionLayout, lights)

	fmt.Println("total1:", total1)

	maxLit := 0

	// right wall going left
	for j := 0; j < len(contraptionLayout); j++ {
		lights := []LightParticle{
			{
				Position: Position{
					X: len(contraptionLayout[0]) - 1,
					Y: j,
				},
				Direction: left,
			},
		}

		lit := calculateLitSquares(contraptionLayout, lights)

		if lit > maxLit {
			maxLit = lit
		}

		fmt.Printf("right wall going left %d/%d %d (%d)\n", j, len(contraptionLayout), lit, maxLit)
	}

	// left wall going right
	for j := 0; j < len(contraptionLayout); j++ {
		lights := []LightParticle{
			{
				Position: Position{
					X: 0,
					Y: j,
				},
				Direction: right,
			},
		}

		lit := calculateLitSquares(contraptionLayout, lights)

		if lit > maxLit {
			maxLit = lit
		}

		fmt.Printf("left wall going right %d/%d %d (%d)\n", j, len(contraptionLayout), lit, maxLit)
	}

	// bottom wall going up
	for i := 0; i < len(contraptionLayout[0]); i++ {
		lights := []LightParticle{
			{
				Position: Position{
					X: i,
					Y: len(contraptionLayout) - 1,
				},
				Direction: up,
			},
		}

		lit := calculateLitSquares(contraptionLayout, lights)

		if lit > maxLit {
			maxLit = lit
		}

		fmt.Printf("right wall going left %d/%d %d (%d)\n", i, len(contraptionLayout), lit, maxLit)
	}

	// top wall going down
	for i := 0; i < len(contraptionLayout[0]); i++ {
		lights := []LightParticle{
			{
				Position: Position{
					X: i,
					Y: 0,
				},
				Direction: down,
			},
		}

		lit := calculateLitSquares(contraptionLayout, lights)

		if lit > maxLit {
			maxLit = lit
		}

		fmt.Printf("right wall going left %d/%d %d (%d)\n", i, len(contraptionLayout), lit, maxLit)
	}

	fmt.Println("maxLit:", maxLit)
}

func calculateLitSquares(contraptionLayout [][]string, startingLights []LightParticle) int {
	lights := startingLights

	litSquares := map[string]bool{}
	everHad := map[string]bool{}
	for {
		newLights := []LightParticle{}

		for _, light := range lights {
			contraptionSquare := contraptionLayout[light.Position.Y][light.Position.X]

			squareAddress := fmt.Sprintf("%d-%d", light.Position.X, light.Position.Y)
			litSquares[squareAddress] = true
			litCount := 0
			for _, isLit := range litSquares {
				if isLit {
					litCount++
				}
			}

			newDirections := getNewDirections(light.Direction, contraptionSquare)

			for _, direction := range newDirections {
				newSquareX := light.Position.X + direction.H
				newSquareY := light.Position.Y + direction.V

				if newSquareX < 0 || newSquareX > (len(contraptionLayout[0])-1) {
					continue
				}

				if newSquareY < 0 || newSquareY > (len(contraptionLayout)-1) {
					continue
				}

				hash := fmt.Sprintf("%d,%d;%d,%d", newSquareX, newSquareY, direction.H, direction.V)

				if everHad[hash] {
					continue
				}

				everHad[hash] = true

				alreadyHave := false
				for _, a := range newLights {
					if a.Hash == hash {
						alreadyHave = true
						break
					}
				}

				if !alreadyHave {
					newLights = append(newLights, LightParticle{
						Position: Position{
							X: newSquareX,
							Y: newSquareY,
						},
						Direction: direction,
						Hash:      hash,
					})
				}
			}
		}

		// for j, row := range contraptionLayout {
		// 	for i, col := range row {
		// 		lightsCount := 0
		// 		light := LightParticle{}
		// 		for _, l := range newLights {
		// 			if i == l.Position.X && j == l.Position.Y {
		// 				lightsCount++
		// 				light = l
		// 			}
		// 		}

		// 		if lightsCount > 1 && lightsCount <= 9 {
		// 			fmt.Printf("%d", lightsCount)
		// 		} else if lightsCount > 9 {
		// 			fmt.Print("9")
		// 		} else if lightsCount == 1 {
		// 			fmt.Print(light.Direction.C)
		// 		} else {
		// 			fmt.Print(col)
		// 		}
		// 	}

		// 	fmt.Println()
		// }

		// fmt.Println()

		if len(newLights) == 0 {
			break
		}

		lights = newLights
	}

	// fmt.Println(contraptionLayout)

	litCount := 0
	for _, isLit := range litSquares {
		if isLit {
			litCount++
		}
	}

	return litCount
}

func getNewDirections(prevDirection Direction, tile string) []Direction {
	switch tile {
	case `\`:
		{
			switch prevDirection {
			case up:
				{
					return []Direction{left}
				}
			case down:
				{
					return []Direction{right}
				}
			case left:
				{
					return []Direction{up}
				}
			case right:
				{
					return []Direction{down}
				}
			}
		}
	case `/`:
		{
			switch prevDirection {
			case up:
				{
					return []Direction{right}
				}
			case down:
				{
					return []Direction{left}
				}
			case left:
				{
					return []Direction{down}
				}
			case right:
				{
					return []Direction{up}
				}
			}
		}
	case `-`:
		{
			switch prevDirection {
			case up:
				fallthrough
			case down:
				{
					return []Direction{left, right}
				}
			case left:
				fallthrough
			case right:
				{
					return []Direction{prevDirection}
				}
			}
		}
	case `|`:
		{
			switch prevDirection {
			case up:
				fallthrough
			case down:
				{
					return []Direction{prevDirection}
				}
			case left:
				fallthrough
			case right:
				{
					return []Direction{up, down}
				}
			}
		}
	}

	return []Direction{prevDirection}
}
