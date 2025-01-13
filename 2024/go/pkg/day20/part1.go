package day20

import (
	"math"
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type coordinate struct {
	x int
	y int
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)
	board := make([][]int, 0, len(input))

	start := coordinate{}
	end := coordinate{}

	for i := range input {
		row := make([]int, len(input[i]))

		for j := range row {
			if string(input[i][j]) == "#" {
				row[j] = 100
				row[j] = math.MaxInt
			}

			if string(input[i][j]) == "S" {
				start.x = i
				start.y = j
			}

			if string(input[i][j]) == "E" {
				end.x = i
				end.y = j
			}
		}

		board = append(board, row)
	}

	position := start

	for position != end {
		// north
		if board[position.x-1][position.y] == 0 && (coordinate{x: position.x - 1, y: position.y} != start) {
			board[position.x-1][position.y] = board[position.x][position.y] + 1
			position.x -= 1
		}

		// south
		if board[position.x+1][position.y] == 0 && (coordinate{x: position.x + 1, y: position.y} != start) {
			board[position.x+1][position.y] = board[position.x][position.y] + 1
			position.x += 1
		}

		// east
		if board[position.x][position.y+1] == 0 && (coordinate{x: position.x, y: position.y + 1} != start) {
			board[position.x][position.y+1] = board[position.x][position.y] + 1
			position.y += 1
		}

		// west
		if board[position.x][position.y-1] == 0 && (coordinate{x: position.x, y: position.y - 1} != start) {
			board[position.x][position.y-1] = board[position.x][position.y] + 1
			position.y -= 1
		}
	}

	result := 0
	position = start

	for position != end {
		next := coordinate{}

		// north
		if board[position.x-1][position.y] == board[position.x][position.y]+1 {
			next.x = position.x - 1
			next.y = position.y
		}

		// south
		if board[position.x+1][position.y] == board[position.x][position.y]+1 {
			next.x = position.x + 1
			next.y = position.y
		}

		// east
		if board[position.x][position.y+1] == board[position.x][position.y]+1 {
			next.x = position.x
			next.y = position.y + 1
		}

		// west
		if board[position.x][position.y-1] == board[position.x][position.y]+1 {
			next.x = position.x
			next.y = position.y - 1
		}

		current := board[position.x][position.y] + 2

		// north
		if board[position.x-1][position.y] == math.MaxInt && position.x-2 > -1 && board[position.x-2][position.y] != math.MaxInt && ((board[position.x][position.y] + 2) < board[position.x-2][position.y]) {
			diff := board[position.x-2][position.y] - current
			if diff >= 100 {
				result += 1
			}
		}

		// south
		if board[position.x+1][position.y] == math.MaxInt && position.x+2 < len(board) && board[position.x+2][position.y] != math.MaxInt && ((board[position.x][position.y] + 2) < board[position.x+2][position.y]) {
			diff := board[position.x+2][position.y] - current
			if diff >= 100 {
				result += 1
			}
		}

		// east
		if board[position.x][position.y+1] == math.MaxInt && position.y+2 < len(board[position.x]) && board[position.x][position.y+2] != math.MaxInt && ((board[position.x][position.y] + 2) < board[position.x][position.y+2]) {
			diff := board[position.x][position.y+2] - current
			if diff >= 100 {
				result += 1
			}
		}

		// west
		if board[position.x][position.y-1] == math.MaxInt && position.y-2 > -1 && board[position.x][position.y-2] != math.MaxInt && ((board[position.x][position.y] + 2) < board[position.x][position.y-2]) {
			diff := board[position.x][position.y-2] - current
			if diff >= 100 {
				result += 1
			}
		}

		position = next
	}

	return result
}
