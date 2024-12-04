package day2

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Execute(part string, file *os.File) error {
	switch part {
	case "1":
		Part1(file)
	case "2":
		Part2(file)
	default:
		log.Fatalf("unknown part: %s", part)
		return errors.New(fmt.Sprintf("unknown part: %s", part))
	}

	return nil
}
