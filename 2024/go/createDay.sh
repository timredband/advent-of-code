#!/usr/bin/env bash

SCRIPT_FOLDER=$(dirname -- "$0")
DAY="$1"

if [[ -z "$DAY" ]]; then
    echo "missing day arg"
    exit 1
fi

if [[ ${#DAY} -eq 1 ]]; then
    DAY="0$DAY"
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

mkdir $SCRIPT_FOLDER/inputs/day$DAY
touch $SCRIPT_FOLDER/inputs/day$DAY/example.txt
touch $SCRIPT_FOLDER/inputs/day$DAY/input.txt
