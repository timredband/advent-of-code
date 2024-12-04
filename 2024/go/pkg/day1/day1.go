package day1

import (
	"errors"
	"fmt"
	"os"
)

func Execute(part string, file *os.File) error {
	switch part {
	case "1":
		Part1(file)
	case "2":
		Part2(file)
	default:
		return errors.New(fmt.Sprintf("unknown part: %s", part))
	}

	return nil
}
