package day22

import (
	"os"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func calculateNextSecret(secret int) int {
	secret = ((secret * 64) ^ secret) % 16777216
	secret = ((secret / 32) ^ secret) % 16777216
	secret = ((secret * 2048) ^ secret) % 16777216

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
