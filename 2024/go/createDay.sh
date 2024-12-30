#!/usr/bin/env bash

SCRIPT_FOLDER=$(dirname -- "$0")
DAY="$1"

if [[ -z "$DAY" ]]; then
    echo "missing day arg"
    exit 1
fi

mkdir $SCRIPT_FOLDER/pkg/day$DAY

cat <<EOF >"$SCRIPT_FOLDER/pkg/day$DAY/part1.go"
package day$DAY

import (
	"os"
)

func Part1(file *os.File) int { return 0 }
EOF

cat <<EOF >"$SCRIPT_FOLDER/pkg/day$DAY/part2.go"
package day$DAY

import (
	"os"
)

func Part2(file *os.File) int { return 0 }
EOF

cat <<EOF >"$SCRIPT_FOLDER/pkg/day$DAY/day$DAY.go"
package day$DAY

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
EOF

mkdir $SCRIPT_FOLDER/inputs/day$DAY
touch $SCRIPT_FOLDER/inputs/day$DAY/example.txt
touch $SCRIPT_FOLDER/inputs/day$DAY/input.txt
