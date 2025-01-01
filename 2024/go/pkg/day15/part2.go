package day15

import (
	"os"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type Range struct {
	low  int
	high int
}

func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	board := make([][]string, 0)
	moves := make([]string, 0)
	position := position{}

	for i := range input {
		if strings.Contains(input[i], "@") {
			position.x = i
			position.y = strings.Index(input[i], "@") * 2
		}

		if !strings.Contains(input[i], "#") {
			moves = append(moves, strings.Split(input[i], "")...)
			continue
		}

		row := make([]string, len(input[i])*2)

		for j := 0; j < len(input[i]); j += 1 {
			col := j * 2

			if string(input[i][j]) == "O" {
				row[col] = "["
				row[col+1] = "]"
				continue
			}

			if string(input[i][j]) == "@" {
				row[col] = "@"
				row[col+1] = "."
				continue
			}

			row[col] = string(input[i][j])
			row[col+1] = string(input[i][j])
		}

		board = append(board, row)
	}

	var boxCanMoveUp func(i int, j int) bool
	boxCanMoveUp = func(i int, j int) bool {
		if board[i-1][j] == "#" || board[i-1][j+1] == "#" {
			return false
		}

		if board[i-1][j] == "." && board[i-1][j+1] == "." {
			return true
		}

		if board[i-1][j] == "[" {
			return boxCanMoveUp(i-1, j)
		}

		left := true
		if board[i-1][j] == "]" {
			left = boxCanMoveUp(i-1, j-1)
		}

		right := true
		if board[i-1][j+1] == "[" {
			right = boxCanMoveUp(i-1, j+1)
		}

		return left && right
	}

	var boxCanMoveDown func(i int, j int) bool
	boxCanMoveDown = func(i int, j int) bool {
		if board[i+1][j] == "#" || board[i+1][j+1] == "#" {
			return false
		}

		if board[i+1][j] == "." && board[i+1][j+1] == "." {
			return true
		}

		if board[i+1][j] == "[" {
			return boxCanMoveDown(i+1, j)
		}

		left := true
		if board[i+1][j] == "]" {
			left = boxCanMoveDown(i+1, j-1)
		}

		right := true
		if board[i+1][j+1] == "[" {
			right = boxCanMoveDown(i+1, j+1)
		}

		return left && right
	}

	var boxMoveUp func(i int, j int)
	boxMoveUp = func(i int, j int) {
		if board[i-1][j] == "[" {
			boxMoveUp(i-1, j)
		}

		if board[i-1][j] == "]" {
			boxMoveUp(i-1, j-1)
		}

		if board[i-1][j+1] == "[" {
			boxMoveUp(i-1, j+1)
		}

		board[i-1][j] = "["
		board[i-1][j+1] = "]"

		board[i][j] = "."
		board[i][j+1] = "."
	}

	var boxMoveDown func(i int, j int)
	boxMoveDown = func(i int, j int) {
		if board[i+1][j] == "[" {
			boxMoveDown(i+1, j)
		}

		if board[i+1][j] == "]" {
			boxMoveDown(i+1, j-1)
		}

		if board[i+1][j+1] == "[" {
			boxMoveDown(i+1, j+1)
		}

		board[i+1][j] = "["
		board[i+1][j+1] = "]"

		board[i][j] = "."
		board[i][j+1] = "."
	}

	for _, move := range moves {
		switch move {
		case "<":
			if board[position.x][position.y-1] == "." {
				board[position.x][position.y] = "."
				board[position.x][position.y-1] = "@"
				position.y -= 1
				continue
			}

			if board[position.x][position.y-1] == "]" {
				x, y := position.x, position.y-1

				for board[x][y] == "]" || board[x][y] == "[" {
					y -= 1
				}

				if board[x][y] == "." {
					for y != position.y {
						board[x][y] = board[x][y+1]
						y += 1
					}
					board[position.x][position.y] = "."
					position.y -= 1
				}
			}
		case ">":
			if board[position.x][position.y+1] == "." {
				board[position.x][position.y] = "."
				board[position.x][position.y+1] = "@"
				position.y += 1
				continue
			}

			if board[position.x][position.y+1] == "[" {
				x, y := position.x, position.y+1

				for board[x][y] == "[" || board[x][y] == "]" {
					y += 1
				}

				if board[x][y] == "." {
					for y != position.y {
						board[x][y] = board[x][y-1]
						y -= 1
					}
					board[position.x][position.y] = "."
					position.y += 1
				}
			}
		case "^":
			if board[position.x-1][position.y] == "#" {
				continue
			}

			canMove := true

			if board[position.x-1][position.y] == "[" {
				canMove = canMove && boxCanMoveUp(position.x-1, position.y)
			}

			if board[position.x-1][position.y] == "]" {
				canMove = canMove && boxCanMoveUp(position.x-1, position.y-1)
			}

			if canMove {
				if board[position.x-1][position.y] == "[" {
					boxMoveUp(position.x-1, position.y)
				}
				if board[position.x-1][position.y] == "]" {
					boxMoveUp(position.x-1, position.y-1)
				}

				board[position.x-1][position.y] = "@"
				board[position.x][position.y] = "."
				position.x -= 1
			}
		case "v":
			if board[position.x+1][position.y] == "#" {
				continue
			}

			canMove := true

			if board[position.x+1][position.y] == "[" {
				canMove = canMove && boxCanMoveDown(position.x+1, position.y)
			}

			if board[position.x+1][position.y] == "]" {
				canMove = canMove && boxCanMoveDown(position.x+1, position.y-1)
			}

			if canMove {
				if board[position.x+1][position.y] == "[" {
					boxMoveDown(position.x+1, position.y)
				}
				if board[position.x+1][position.y] == "]" {
					boxMoveDown(position.x+1, position.y-1)
				}

				board[position.x+1][position.y] = "@"
				board[position.x][position.y] = "."
				position.x += 1
			}
		}
	}

	result := 0

	for i := range board {
		for j := range board[i] {
			if board[i][j] == "[" {
				result += ((100 * i) + j)
			}
		}
	}

	return result
}
