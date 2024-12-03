package day1

import (
	"log"
	"os"
)

func Execute(part string, file *os.File) {
	switch part {
	case "1":
		Part1(file)
	case "2":
		Part2(file)
	default:
		log.Fatalf("unknown part: %s", part)
	}
}
