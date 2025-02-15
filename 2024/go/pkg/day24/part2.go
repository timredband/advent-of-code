package day24

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

// https://edotor.net/ for visualizations
func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	wiresByName, instructions := initialize(input)

	a := 0
	o := 0
	xor := 0

	for _, instruction := range instructions {
		switch instruction.gateType {
		case "AND":
			andString := strconv.Itoa(a)
			fmt.Printf("node [shape = house]; AND%s [label = \"AND\"];\n", andString)
			fmt.Printf("node [shape = circle]; %s\n", instruction.leftInputWire)
			fmt.Printf("node [shape = circle]; %s\n", instruction.rightInputWire)
			fmt.Printf("%s -> AND%s \n", instruction.leftInputWire, andString)
			fmt.Printf("%s -> AND%s \n", instruction.rightInputWire, andString)
			fmt.Printf("AND%s -> %s \n", andString, instruction.destinationWire)
			a += 1
		case "OR":
			orString := strconv.Itoa(o)
			fmt.Printf("node [shape = invtriangle]; OR%s [label = \"OR\"];\n", orString)
			fmt.Printf("node [shape = circle]; %s\n", instruction.leftInputWire)
			fmt.Printf("node [shape = circle]; %s\n", instruction.rightInputWire)
			fmt.Printf("%s -> OR%s \n", instruction.leftInputWire, orString)
			fmt.Printf("%s -> OR%s \n", instruction.rightInputWire, orString)
			fmt.Printf("OR%s -> %s \n", orString, instruction.destinationWire)
			o += 1
		case "XOR":
			xorString := strconv.Itoa(xor)
			fmt.Printf("node [shape = doublecircle]; XOR%s [label = \"XOR\"];\n", xorString)
			fmt.Printf("node [shape = circle]; %s\n", instruction.leftInputWire)
			fmt.Printf("node [shape = circle]; %s\n", instruction.rightInputWire)
			fmt.Printf("%s -> XOR%s \n", instruction.leftInputWire, xorString)
			fmt.Printf("%s -> XOR%s \n", instruction.rightInputWire, xorString)
			fmt.Printf("XOR%s -> %s \n", xorString, instruction.destinationWire)
			xor += 1
		}
	}

	// fvw,grf,nwg,mdb,wpq,z18,z22,z36

	processInstructions(wiresByName, instructions)

	i := 0
	zValue := 0

	for {
		key := fmt.Sprintf("z%02d", i)

		if value, ok := wiresByName[key]; ok {
			zValue += (value * int(math.Pow(2, float64(i))))
			i += 1
		} else {
			break
		}
	}

	// x and y values
	i = 0
	xValue := 0
	yValue := 0

	for {
		xKey := fmt.Sprintf("x%02d", i)
		yKey := fmt.Sprintf("y%02d", i)

		if value, ok := wiresByName[xKey]; ok {
			xValue += (value * int(math.Pow(2, float64(i))))
		}

		if value, ok := wiresByName[yKey]; ok {
			yValue += (value * int(math.Pow(2, float64(i))))
		} else {
			break
		}

		i += 1
	}

	fmt.Printf("xValue: %046b, %d\n", xValue, xValue)
	fmt.Printf("yValue: %046b, %d\n", yValue, yValue)
	fmt.Printf("+Value: %046b, %d\n", xValue+yValue, xValue+yValue)
	fmt.Printf("zValue: %046b, %d\n", zValue, zValue)

	resultValue := fmt.Sprintf("%046b", xValue+yValue)
	expectedValue := fmt.Sprintf("%046b", zValue)

	j := 0
	for i := len(resultValue) - 1; i > -1; i -= 1 {
		if string(resultValue[i]) != string(expectedValue[i]) {
			fmt.Printf("z%02d\n", j)
		}
		j += 1
	}

	fmt.Println("fvw,grf,mdb,nwq,wpq,z18,z22,z36")

	return 0
}
