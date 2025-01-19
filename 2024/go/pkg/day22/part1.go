package day22

import (
	"os"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func calculateNextSecret(secret int) int {
	mask := 0b1111_1111_1111_1111_1111_1111 // 24 bits
	secret = ((secret << 6) ^ secret) & mask
	secret = ((secret >> 5) ^ secret) & mask
	secret = ((secret << 11) ^ secret) & mask

	return secret
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)

	result := 0

	for i := range input {
		secret, _ := strconv.Atoi(input[i])

		for range 2000 {
			secret = calculateNextSecret(secret)
		}

		result += secret
	}

	return result
}
