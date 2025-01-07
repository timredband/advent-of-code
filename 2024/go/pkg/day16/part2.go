package day16

import (
	"math"
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type coordinate struct {
	x int
	y int
}

func Part2(file *os.File) int {
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

	results := make(map[coordinate]bool)

	var checkDFS func(i int, j int, cost int, direction string)
	checkDFS = func(i int, j int, cost int, direction string) {
		if i == startX && j == startY {
			results[coordinate{x: i, y: j}] = true
			return
		}

		canGoAtLeastOneDirection := false

		if visited[i-1][j] < cost {
			canGoAtLeastOneDirection = true
		}
		if visited[i+1][j] < cost {
			canGoAtLeastOneDirection = true
		}
		if visited[i][j+1] < cost {
			canGoAtLeastOneDirection = true
		}
		if visited[i][j-1] < cost {
			canGoAtLeastOneDirection = true
		}

		if !canGoAtLeastOneDirection {
			return
		}

		results[coordinate{x: i, y: j}] = true

		// north
		if visited[i-1][j] < cost {
			c := cost - 1
			if direction != "north" {
				c -= 1000
			}
			checkDFS(i-1, j, c, "north")
		}

		// south
		if visited[i+1][j] < cost {
			c := cost - 1
			if direction != "south" {
				c -= 1000
			}
			checkDFS(i+1, j, c, "south")
		}

		// east
		if visited[i][j+1] < cost {
			c := cost - 1
			if direction != "east" {
				c -= 1000
			}
			checkDFS(i, j+1, c, "east")
		}

		// west
		if visited[i][j-1] < cost {
			c := cost - 1
			if direction != "west" {
				c -= 1000
			}
			checkDFS(i, j-1, c, "west")
		}
	}

	checkDFS(endX, endY, visited[endX][endY], "north")
	checkDFS(endX, endY, visited[endX][endY], "south")
	checkDFS(endX, endY, visited[endX][endY], "east")
	checkDFS(endX, endY, visited[endX][endY], "west")

	return len(results)
}
