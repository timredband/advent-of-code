package day16

import (
	"math"
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
	input := utils.ReadFile(file)
	board := make([][]string, 0, len(input))
	visited := make([][]int, 0, len(board))

	startX, startY := 0, 0
	endX, endY := 0, 0

	for i := range input {
		row := make([]string, len(input[i]))
		visitedRow := make([]int, len(input[i]))

		for j := range input[i] {
			row[j] = string(input[i][j])
			visitedRow[j] = math.MaxInt

			if row[j] == "S" {
				startX, startY = i, j
			}

			if row[j] == "E" {
				endX, endY = i, j
			}
		}

		board = append(board, row)
		visited = append(visited, visitedRow)
	}

	var DFS func(i int, j int, cost int, direction string)
	DFS = func(i int, j int, cost int, direction string) {
		if board[i][j] == "#" {
			return
		}

		if visited[i][j] < cost {
			return
		}

		if board[i][j] == "E" {
			visited[i][j] = min(visited[i][j], cost)
			// more logic here?
			return
		}

		visited[i][j] = cost

		d := "north"
		c := cost + 1
		if d != direction {
			c += 1000
		}
		DFS(i-1, j, c, d)

		d = "south"
		c = cost + 1
		if d != direction {
			c += 1000
		}
		DFS(i+1, j, c, d)

		d = "east"
		c = cost + 1
		if d != direction {
			c += 1000
		}
		DFS(i, j+1, c, d)

		d = "west"
		c = cost + 1
		if d != direction {
			c += 1000
		}
		DFS(i, j-1, c, d)
	}

	DFS(startX, startY, 0, "east")

	return visited[endX][endY]
}
