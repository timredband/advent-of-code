package day10

import (
	"os"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type Coordinate struct {
	x int
	y int
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)
	board := make([][]int, 0, len(input))
	destinationsBoard := make([][]map[Coordinate]struct{}, 0, len(input))

	for i := range input {
		row := make([]int, len(input[i]))
		destinationRow := make([]map[Coordinate]struct{}, len(input[i]))

		for j := range input[i] {
			row[j], _ = strconv.Atoi(string(input[i][j]))
			destinationRow[j] = make(map[Coordinate]struct{})

			if row[j] == 9 {
				destinationRow[j][Coordinate{x: i, y: j}] = struct{}{}
			}
		}

		board = append(board, row)
		destinationsBoard = append(destinationsBoard, destinationRow)
	}

	for score := 8; score > -1; score -= 1 {
		for i := range board {
			for j := range board[i] {
				if board[i][j] == score {
					// up
					if i-1 > -1 && board[i-1][j] == score+1 {
						for k := range destinationsBoard[i-1][j] {
							destinationsBoard[i][j][k] = struct{}{}
						}
					}

					// down
					if i+1 < len(board) && board[i+1][j] == score+1 {
						for k := range destinationsBoard[i+1][j] {
							destinationsBoard[i][j][k] = struct{}{}
						}
					}

					// left
					if j-1 > -1 && board[i][j-1] == score+1 {
						for k := range destinationsBoard[i][j-1] {
							destinationsBoard[i][j][k] = struct{}{}
						}
					}

					// right
					if j+1 < len(board[i]) && board[i][j+1] == score+1 {
						for k := range destinationsBoard[i][j+1] {
							destinationsBoard[i][j][k] = struct{}{}
						}
					}
				}
			}
		}

	}

	result := 0

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				result += (len(destinationsBoard[i][j]))
			}
		}
	}

	return result
}
