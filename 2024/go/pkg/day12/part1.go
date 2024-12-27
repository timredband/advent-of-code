package day12

import (
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
	input := utils.ReadFile(file)
	board := make([][]string, len(input))
	visited := make([][]bool, len(input))

	for i := range input {
		board[i] = make([]string, len(input[i]))
		visited[i] = make([]bool, len(input[i]))

		for j := range input[i] {
			board[i][j] = string(input[i][j])
		}
	}

	var DFS func(i int, j int, current string) (area int, perimeter int)
	DFS = func(i int, j int, current string) (area int, perimeter int) {
		if visited[i][j] {
			return
		}

		visited[i][j] = true

		area += 1

		// up
		if i-1 > -1 && board[i-1][j] == current {
			upArea, upPerimeter := DFS(i-1, j, current)
			area += upArea
			perimeter += upPerimeter
		} else {
			perimeter += 1
		}

		// down
		if i+1 < len(board) && board[i+1][j] == current {
			downArea, downPerimeter := DFS(i+1, j, current)
			area += downArea
			perimeter += downPerimeter
		} else {
			perimeter += 1
		}

		// left
		if j-1 > -1 && board[i][j-1] == current {
			leftArea, leftPerimeter := DFS(i, j-1, current)
			area += leftArea
			perimeter += leftPerimeter
		} else {
			perimeter += 1
		}

		// right
		if j+1 < len(board[i]) && board[i][j+1] == current {
			rightArea, rightPerimeter := DFS(i, j+1, current)
			area += rightArea
			perimeter += rightPerimeter
		} else {
			perimeter += 1
		}

		return
	}

	result := 0
	for i := range board {
		for j := range board[i] {
			area, perimeter := DFS(i, j, board[i][j])
			result += (area * perimeter)
		}
	}

	return result
}
