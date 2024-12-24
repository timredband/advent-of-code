package day10

import (
	"os"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	input := utils.ReadFile(file)
	board := make([][]int, 0, len(input))
	DP := make([][]int, len(input))

	for i := range input {
		row := make([]int, len(input[i]))
		DP[i] = make([]int, len(input[i]))

		for j := range input[i] {
			if string(input[i][j]) == "." {
				row[j] = -1
			} else {
				row[j], _ = strconv.Atoi(string(input[i][j]))
			}

			if row[j] == 9 {
				DP[i][j] = 1
			}
		}

		board = append(board, row)
	}

	for score := 8; score > -1; score -= 1 {
		for i := range board {
			for j := range board[i] {
				if board[i][j] == score {
					// up
					if i-1 > -1 && board[i-1][j] == score+1 {
						DP[i][j] += DP[i-1][j]
					}

					// down
					if i+1 < len(board) && board[i+1][j] == score+1 {
						DP[i][j] += DP[i+1][j]
					}

					// left
					if j-1 > -1 && board[i][j-1] == score+1 {
						DP[i][j] += DP[i][j-1]
					}

					// right
					if j+1 < len(board[i]) && board[i][j+1] == score+1 {
						DP[i][j] += DP[i][j+1]
					}
				}
			}
		}
	}

	result := 0

	for i := range DP {
		for j := range DP[i] {
			if board[i][j] == 0 {
				result += DP[i][j]
			}
		}
	}

	return result
}
