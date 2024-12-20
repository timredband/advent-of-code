package day6

import (
	"fmt"
	"os"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type Position struct {
	i         int
	j         int
	direction string
}

func Part2(file *os.File) int {
	inputs := utils.ReadFile(file)

	initialBoard := make([][]string, len(inputs))

	initialPosition := Position{direction: "north"}
	// maxPosition := len(inputs) - 1

	for i, v := range inputs {
		initialBoard[i] = strings.Split(v, "")

		for j, c := range initialBoard[i] {
			if c == "^" {
				initialPosition.i = i
				initialPosition.j = j
			}
		}
	}

	for i := range initialBoard {
		for j := range initialBoard[i] {
			board := make([][]string, len(initialBoard))
			for k := range initialBoard {
				board[k] = make([]string, len(initialBoard[k]))
				copy(board[k], initialBoard[k])
			}
		}
	}

	for i := range board {
		fmt.Println(board[i])
	}

	// for position.i > 0 && position.j < len(inputs) {
	// 	board[position.i][position.j] = "X"
	//
	// 	if position.direction == "north" {
	// 		if position.i-1 < 0 {
	// 			break
	// 		}
	//
	// 		if board[position.i-1][position.j] == "#" {
	// 			position.direction = "east"
	// 		} else {
	// 			position.i -= 1
	// 		}
	//
	// 		continue
	// 	}
	//
	// 	if position.direction == "south" {
	// 		if position.i+1 > maxPosition {
	// 			break
	// 		}
	//
	// 		if board[position.i+1][position.j] == "#" {
	// 			position.direction = "west"
	// 		} else {
	// 			position.i += 1
	// 		}
	//
	// 		continue
	// 	}
	//
	// 	if position.direction == "east" {
	// 		if position.j+1 > maxPosition {
	// 			break
	// 		}
	//
	// 		if board[position.i][position.j+1] == "#" {
	// 			position.direction = "south"
	// 		} else {
	// 			position.j += 1
	// 		}
	//
	// 		continue
	//
	// 	}
	//
	// 	if position.direction == "west" {
	// 		if position.j-1 < 0 {
	// 			break
	// 		}
	//
	// 		if board[position.i][position.j-1] == "#" {
	// 			position.direction = "north"
	// 		} else {
	// 			position.j -= 1
	// 		}
	//
	// 		continue
	// 	}
	// }
	//
	// result := 0
	//
	// for _, row := range board {
	// 	for _, v := range row {
	// 		if v == "X" {
	// 			result += 1
	// 		}
	// 	}
	// }

	return 0
}
