package day08

import (
	"os"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type Coordinate struct {
	x int
	y int
}

func Part1(file *os.File) int {
	inputs := utils.ReadFile(file)
	board := make([][]string, len(inputs))

	for i, v := range inputs {
		line := strings.Split(v, "")
		board[i] = line
	}

	coordinatesByAntenna := make(map[string][]Coordinate)

	for x, row := range board {
		for y, val := range row {
			if val == "." {
				continue
			}

			if _, ok := coordinatesByAntenna[val]; !ok {
				coordinatesByAntenna[val] = make([]Coordinate, 0)
			}

			coordinatesByAntenna[val] = append(coordinatesByAntenna[val], Coordinate{x, y})
		}

	}

	antinodeCoordinates := make(map[Coordinate]struct{})

	for _, coordinates := range coordinatesByAntenna {
		i, j := 0, 1

		for i < len(coordinates) {
			for j < len(coordinates) {
				xDiff := coordinates[i].x - coordinates[j].x
				yDiff := coordinates[i].y - coordinates[j].y

				firstAntiNode := Coordinate{x: coordinates[i].x + xDiff, y: coordinates[i].y + yDiff}
				secondAntiNode := Coordinate{x: coordinates[j].x - xDiff, y: coordinates[j].y - yDiff}

				if firstAntiNode.x >= 0 && firstAntiNode.x < len(board) && firstAntiNode.y >= 0 && firstAntiNode.y < len(board) {
					antinodeCoordinates[firstAntiNode] = struct{}{}
				}

				if secondAntiNode.x >= 0 && secondAntiNode.x < len(board) && secondAntiNode.y >= 0 && secondAntiNode.y < len(board) {
					antinodeCoordinates[secondAntiNode] = struct{}{}
				}

				j += 1
			}

			coordinates = coordinates[1:]

			i, j = 0, 1
		}
	}

	result := len(antinodeCoordinates)

	return result
}
