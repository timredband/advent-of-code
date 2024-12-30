package day15

import (
	"os"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type position struct {
	x int
	y int
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)

	board := make([][]string, 0)
	moves := make([]string, 0)
	position := position{}

	for i := range input {
		if strings.Contains(input[i], "@") {
			position.x = i
			position.y = strings.Index(input[i], "@")
		}

		if !strings.Contains(input[i], "#") {
			moves = append(moves, strings.Split(input[i], "")...)
			continue
		}

		row := make([]string, len(input[i]))

		for j := range input[i] {
			row[j] = string(input[i][j])
		}

		board = append(board, row)
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

			if board[position.x][position.y-1] == "O" {
				x, y := position.x, position.y-1

				for board[x][y] == "O" {
					y -= 1
				}

				if board[x][y] == "." {
					board[x][y] = "O"
					board[position.x][position.y-1] = "@"
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

			if board[position.x][position.y+1] == "O" {
				x, y := position.x, position.y+1

				for board[x][y] == "O" {
					y += 1
				}

				if board[x][y] == "." {
					board[x][y] = "O"
					board[position.x][position.y+1] = "@"
					board[position.x][position.y] = "."
					position.y += 1
				}
			}
		case "^":
			if board[position.x-1][position.y] == "." {
				board[position.x][position.y] = "."
				board[position.x-1][position.y] = "@"
				position.x -= 1
				continue
			}

			if board[position.x-1][position.y] == "O" {
				x, y := position.x-1, position.y

				for board[x][y] == "O" {
					x -= 1
				}

				if board[x][y] == "." {
					board[x][y] = "O"
					board[position.x-1][position.y] = "@"
					board[position.x][position.y] = "."
					position.x -= 1
				}
			}
		case "v":
			if board[position.x+1][position.y] == "." {
				board[position.x][position.y] = "."
				board[position.x+1][position.y] = "@"
				position.x += 1
				continue
			}

			if board[position.x+1][position.y] == "O" {
				x, y := position.x+1, position.y

				for board[x][y] == "O" {
					x += 1
				}

				if board[x][y] == "." {
					board[x][y] = "O"
					board[position.x+1][position.y] = "@"
					board[position.x][position.y] = "."
					position.x += 1
				}
			}
		}
	}

	result := 0

	for i := range board {
		for j := range board[i] {
			if board[i][j] == "O" {
				result += ((100 * i) + j)
			}
		}
	}

	return result
}
