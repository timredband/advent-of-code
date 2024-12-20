package day2

import (
	"errors"
	"fmt"
	"os"
)

func Execute(part string, file *os.File) (int, error) {
	switch part {
	case "1":
		return Part1(file), nil
	case "2":
		return Part2(file), nil
	default:
		return 0, errors.New(fmt.Sprintf("unknown part: %s", part))
	}
}
