package day8

import (
	"os"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
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
		antennaIsAntiNode := false

		if len(coordinates) > 1 {
			antennaIsAntiNode = true
		}

		i, j := 0, 1

		for i < len(coordinates) {
			if antennaIsAntiNode {
				antinodeCoordinates[coordinates[i]] = struct{}{}
			}

			for j < len(coordinates) {
				xDiff := coordinates[i].x - coordinates[j].x
				yDiff := coordinates[i].y - coordinates[j].y

				firstAntiNode := Coordinate{x: coordinates[i].x + xDiff, y: coordinates[i].y + yDiff}
				secondAntiNode := Coordinate{x: coordinates[j].x - xDiff, y: coordinates[j].y - yDiff}

				for firstAntiNode.x >= 0 && firstAntiNode.x < len(board) && firstAntiNode.y >= 0 && firstAntiNode.y < len(board) {
					firstCopy := firstAntiNode
					antinodeCoordinates[firstCopy] = struct{}{}
					firstAntiNode.x += xDiff
					firstAntiNode.y += yDiff
				}

				for secondAntiNode.x >= 0 && secondAntiNode.x < len(board) && secondAntiNode.y >= 0 && secondAntiNode.y < len(board) {
					secondCopy := secondAntiNode
					antinodeCoordinates[secondCopy] = struct{}{}
					secondAntiNode.x -= xDiff
					secondAntiNode.y -= yDiff
				}

				j += 1
			}

			coordinates = coordinates[1:]

			i, j = 0, 1
		}
	}

	for i := range antinodeCoordinates {
		board[i.x][i.y] = "#"
	}

	result := len(antinodeCoordinates)

	return result
}
