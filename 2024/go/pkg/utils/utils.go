package utils

import (
	"bufio"
	"os"
)

func ReadFile(file *os.File) []string {
	var inputs []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		inputs = append(inputs, line)
	}

	return inputs
}
