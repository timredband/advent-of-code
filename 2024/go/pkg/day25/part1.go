package day25

import (
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
	inputs := utils.ReadFile(file)
	locks := make([][]int, 0)
	keys := make([][]int, 0)

	for i := 0; i < len(inputs); i += 7 {
		// lock
		if string(inputs[i][0]) == "#" {
			lock := make([]int, 0, 5)

			for j := 0; j < 5; j += 1 {
				for k := i; k < i+7; k += 1 {
					if string(inputs[k][j]) == "." {
						lock = append(lock, k-i-1)
						break
					}
				}
			}
			locks = append(locks, lock)
		}

		// key
		if string(inputs[i][0]) == "." {
			key := make([]int, 0, 5)

			for j := 0; j < 5; j += 1 {
				for k := i; k < i+7; k += 1 {
					if string(inputs[k][j]) == "#" {
						key = append(key, 6-(k-i))
						break
					}
				}
			}

			keys = append(keys, key)
		}

	}

	result := 0

	for _, key := range keys {
	locks:
		for _, lock := range locks {
			for i := range 5 {
				if key[i]+lock[i] > 5 {
					continue locks
				}
			}

			result += 1
		}
	}

	return result
}
