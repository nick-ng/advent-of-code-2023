package day10

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"
)

type coordinateType struct {
	Row    int
	Column int
}

type pipeType struct {
	Character         string
	Position          coordinateType
	PositionString    string
	Connections       []coordinateType
	IsStart           bool
	IsGround          bool
	DistanceFromStart int
}

func Run() {
	fmt.Println("day 10")

	rawData := utils.ReadFile("10/data.txt")

	pipes := map[string]pipeType{}

	rows := strings.Split(rawData, "\n")

	startPipe := pipeType{}

	pipeCount := 0

	for i, row := range rows {
		if len(row) == 0 {
			continue
		}

		columns := strings.Split(row, "")

		for j, pipeCharacter := range columns {
			if len(pipeCharacter) == 0 {
				continue
			}

			coordinate := coordinateType{
				Row:    i,
				Column: j,
			}

			if pipeCharacter != "." {
				pipeCount++
			}

			pipe := getPipe(pipeCharacter, coordinate)

			pipes[pipe.PositionString] = pipe

			if pipe.IsStart {
				startPipe = pipe
			}
		}
	}

	fmt.Println("pipeCount", pipeCount)

	// Find out the shape of the start pipe
	for _, pipe := range pipes {
		if len(pipe.Connections) == 0 {
			continue
		}

		isConnectedToStart := false

		for _, connection := range pipe.Connections {
			if connection == startPipe.Position {
				isConnectedToStart = true
			}
		}

		if isConnectedToStart {
			startPipe.Connections = append(startPipe.Connections, pipe.Position)
		}
	}

	keepGoing := true
	forwardDistance := 0
	prevForwardPipe := pipeType{}
	forwardPipe := startPipe

	visitedPipes := []string{}
	leftSquares := map[string]bool{}
	rightSquares := map[string]bool{}
	turns := map[string]int{
		"left":     0,
		"right":    0,
		"straight": 0,
	}

	for keepGoing {
		if forwardPipe.IsStart {
			nextConnection := forwardPipe.Connections[0]
			forwardConnectionString := fmt.Sprintf("%d-%d", nextConnection.Row, nextConnection.Column)

			prevForwardPipe = forwardPipe
			forwardPipe = pipes[forwardConnectionString]
			visitedPipes = append(visitedPipes, forwardConnectionString)

			forwardDistance++
			forwardPipe.DistanceFromStart = forwardDistance
			pipes[forwardConnectionString] = forwardPipe

			left, right := getAdjacentPipes(prevForwardPipe.Position, forwardPipe.Position)

			for _, cs := range left {
				leftSquares[cs] = true
			}

			for _, cs := range right {
				rightSquares[cs] = true
			}

			turn := getTurnDirection(prevForwardPipe.Position, forwardPipe.Position, forwardPipe.Character)

			turns[turn]++
		} else {
			// forwardPipe
			forwardConnectionString := ""
			for _, connection := range forwardPipe.Connections {
				if connection != prevForwardPipe.Position {
					forwardConnectionString = fmt.Sprintf("%d-%d", connection.Row, connection.Column)
				}
			}

			prevForwardPipe = forwardPipe
			forwardPipe = pipes[forwardConnectionString]
			visitedPipes = append(visitedPipes, forwardConnectionString)
			forwardDistance++
			forwardPipe.DistanceFromStart = forwardDistance
			pipes[forwardConnectionString] = forwardPipe

			left, right := getAdjacentPipes(prevForwardPipe.Position, forwardPipe.Position)

			for _, cs := range left {
				leftSquares[cs] = true
			}

			for _, cs := range right {
				rightSquares[cs] = true
			}

			turn := getTurnDirection(prevForwardPipe.Position, forwardPipe.Position, forwardPipe.Character)

			turns[turn]++

			if forwardPipe.IsStart {
				keepGoing = false
			}

			if len(visitedPipes) > pipeCount {
				fmt.Println("visited too many pipes")
				// os.Exit(1)
			}
		}
	}

	fmt.Println("farthest", forwardDistance/2)

	insideSquaresTemp := map[string]bool{}

	if turns["left"] > turns["right"] {
		insideSquaresTemp = leftSquares
	} else {
		insideSquaresTemp = rightSquares
	}

	nextSquares := []string{}

	for coordinateString := range insideSquaresTemp {
		nextSquares = append(nextSquares, coordinateString)
	}

	adjacent := []int{-1, 1}
	insideSquares := map[string]bool{}
	checkedSquares := []string{}
	keepGoing = true
	for keepGoing {
		currentSquares := make([]string, len(nextSquares))
		copy(currentSquares, nextSquares)
		nextSquares = []string{}

		keepGoing = false
		for _, coordStr := range currentSquares {
			checkedSquares = append(checkedSquares, coordStr)
			currentSquare := pipes[coordStr]
			if currentSquare.DistanceFromStart < 0 {
				insideSquares[coordStr] = true

				for _, i := range adjacent {
					for _, j := range adjacent {
						coordStr2 := fmt.Sprintf("%d-%d", currentSquare.Position.Row+i, currentSquare.Position.Column+j)

						nextSquare := pipes[coordStr2]

						if utils.SliceContainsString(currentSquares, coordStr2) || utils.SliceContainsString(nextSquares, coordStr2) || utils.SliceContainsString(checkedSquares, coordStr2) {
							continue
						}

						if nextSquare.DistanceFromStart < 0 {
							nextSquares = append(nextSquares, coordStr2)
							keepGoing = true
						}
					}
				}
			}
		}
	}

	fmt.Println("insideSquares", len(insideSquares))
}

func getPipe(pipeCharacter string, coordinate coordinateType) pipeType {
	switch pipeCharacter {
	case "|":
		{
			return pipeType{
				Character:      pipeCharacter,
				Position:       coordinate,
				PositionString: fmt.Sprintf("%d-%d", coordinate.Row, coordinate.Column),
				Connections: []coordinateType{
					{
						Row:    coordinate.Row - 1,
						Column: coordinate.Column,
					},
					{
						Row:    coordinate.Row + 1,
						Column: coordinate.Column,
					},
				},
				IsStart:           false,
				IsGround:          false,
				DistanceFromStart: -1,
			}
		}
	case "-":
		{
			return pipeType{
				Character:      pipeCharacter,
				Position:       coordinate,
				PositionString: fmt.Sprintf("%d-%d", coordinate.Row, coordinate.Column),
				Connections: []coordinateType{
					{
						Row:    coordinate.Row,
						Column: coordinate.Column - 1,
					},
					{
						Row:    coordinate.Row,
						Column: coordinate.Column + 1,
					},
				},
				IsStart:           false,
				IsGround:          false,
				DistanceFromStart: -1,
			}
		}
	case "L":
		{
			return pipeType{
				Character:      pipeCharacter,
				Position:       coordinate,
				PositionString: fmt.Sprintf("%d-%d", coordinate.Row, coordinate.Column),
				Connections: []coordinateType{
					{
						Row:    coordinate.Row,
						Column: coordinate.Column + 1,
					},
					{
						Row:    coordinate.Row - 1,
						Column: coordinate.Column,
					},
				},
				IsStart:           false,
				IsGround:          false,
				DistanceFromStart: -1,
			}
		}
	case "J":
		{
			return pipeType{
				Character:      pipeCharacter,
				Position:       coordinate,
				PositionString: fmt.Sprintf("%d-%d", coordinate.Row, coordinate.Column),
				Connections: []coordinateType{
					{
						Row:    coordinate.Row,
						Column: coordinate.Column - 1,
					},
					{
						Row:    coordinate.Row - 1,
						Column: coordinate.Column,
					},
				},
				IsStart:           false,
				IsGround:          false,
				DistanceFromStart: -1,
			}
		}
	case "7":
		{
			return pipeType{
				Character:      pipeCharacter,
				Position:       coordinate,
				PositionString: fmt.Sprintf("%d-%d", coordinate.Row, coordinate.Column),
				Connections: []coordinateType{
					{
						Row:    coordinate.Row,
						Column: coordinate.Column - 1,
					},
					{
						Row:    coordinate.Row + 1,
						Column: coordinate.Column,
					},
				},
				IsStart:           false,
				IsGround:          false,
				DistanceFromStart: -1,
			}
		}
	case "F":
		{
			return pipeType{
				Character:      pipeCharacter,
				Position:       coordinate,
				PositionString: fmt.Sprintf("%d-%d", coordinate.Row, coordinate.Column),
				Connections: []coordinateType{
					{
						Row:    coordinate.Row,
						Column: coordinate.Column + 1,
					},
					{
						Row:    coordinate.Row + 1,
						Column: coordinate.Column,
					},
				},
				IsStart:           false,
				IsGround:          false,
				DistanceFromStart: -1,
			}
		}
	case "S":
		{
			return pipeType{
				Character:         pipeCharacter,
				Position:          coordinate,
				PositionString:    fmt.Sprintf("%d-%d", coordinate.Row, coordinate.Column),
				Connections:       []coordinateType{},
				IsStart:           true,
				IsGround:          false,
				DistanceFromStart: 0,
			}
		}
	default: // Ground
		{
			return pipeType{
				Character:         pipeCharacter,
				Position:          coordinate,
				PositionString:    fmt.Sprintf("%d-%d", coordinate.Row, coordinate.Column),
				Connections:       []coordinateType{},
				IsStart:           false,
				IsGround:          true,
				DistanceFromStart: -1,
			}
		}

	}
}

func getAdjacentPipes(coord1, coord2 coordinateType) ([]string, []string) {
	if coord1.Column == coord2.Column && coord1.Row < coord2.Row {
		// Down
		l1 := fmt.Sprintf("%d-%d", coord1.Row, coord1.Column+1)
		l2 := fmt.Sprintf("%d-%d", coord2.Row, coord2.Column+1)
		left := []string{l1, l2}

		r1 := fmt.Sprintf("%d-%d", coord1.Row, coord1.Column-1)
		r2 := fmt.Sprintf("%d-%d", coord2.Row, coord2.Column-1)
		right := []string{r1, r2}

		return left, right
	} else if coord1.Column == coord2.Column && coord1.Row > coord2.Row {
		// Up
		l1 := fmt.Sprintf("%d-%d", coord1.Row, coord1.Column-1)
		l2 := fmt.Sprintf("%d-%d", coord2.Row, coord2.Column-1)
		left := []string{l1, l2}

		r1 := fmt.Sprintf("%d-%d", coord1.Row, coord1.Column+1)
		r2 := fmt.Sprintf("%d-%d", coord2.Row, coord2.Column+1)
		right := []string{r1, r2}

		return left, right
	} else if coord1.Row == coord2.Row && coord1.Column < coord2.Column {
		// Right
		l1 := fmt.Sprintf("%d-%d", coord1.Row-1, coord1.Column)
		l2 := fmt.Sprintf("%d-%d", coord2.Row-1, coord2.Column)
		left := []string{l1, l2}

		r1 := fmt.Sprintf("%d-%d", coord1.Row+1, coord1.Column)
		r2 := fmt.Sprintf("%d-%d", coord2.Row+1, coord2.Column)
		right := []string{r1, r2}

		return left, right
	} else if coord1.Row == coord2.Row && coord1.Column > coord2.Column {
		// Left
		l1 := fmt.Sprintf("%d-%d", coord1.Row+1, coord1.Column)
		l2 := fmt.Sprintf("%d-%d", coord2.Row+1, coord2.Column)
		left := []string{l1, l2}

		r1 := fmt.Sprintf("%d-%d", coord1.Row-1, coord1.Column)
		r2 := fmt.Sprintf("%d-%d", coord2.Row-1, coord2.Column)
		right := []string{r1, r2}

		return left, right
	}

	fmt.Println("something went wrong when getting adjacent pipes", coord1, coord2)
	return nil, nil
}

func getTurnDirection(coord1, coord2 coordinateType, destinationCharacter string) string {
	if coord1.Column == coord2.Column && coord1.Row < coord2.Row {
		// Down
		switch destinationCharacter {
		case "L":
			{
				return "left"
			}
		case "J":
			{
				return "right"
			}
		}
	} else if coord1.Column == coord2.Column && coord1.Row > coord2.Row {
		// Up
		switch destinationCharacter {
		case "7":
			{
				return "left"
			}
		case "F":
			{
				return "right"
			}
		}
	} else if coord1.Row == coord2.Row && coord1.Column < coord2.Column {
		// Right
		switch destinationCharacter {
		case "J":
			{
				return "left"
			}
		case "7":
			{
				return "right"
			}
		}
	} else if coord1.Row == coord2.Row && coord1.Column > coord2.Column {
		// Left
		switch destinationCharacter {
		case "F":
			{
				return "left"
			}
		case "L":
			{
				return "right"
			}
		}
	}
	return "straight"
}

func getCoordinateString(coord coordinateType) string {
	coordinateString := fmt.Sprintf("%d-%d", coord.Row, coord.Column)
	return coordinateString
}
