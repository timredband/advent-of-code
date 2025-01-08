package day17

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

// this is not a solution. TODO: figure this out
func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	r := regexp.MustCompile("[0-9]+")
	registerA, _ := strconv.Atoi(r.FindAllString(input[0], -1)[0])
	registerB, _ := strconv.Atoi(r.FindAllString(input[1], -1)[0])
	registerC, _ := strconv.Atoi(r.FindAllString(input[2], -1)[0])

	program := make([]int, 0)

	nums := r.FindAllString(input[3], -1)
	for i := range nums {
		num, _ := strconv.Atoi(nums[i])
		program = append(program, num)
	}

	var getComboOperand func(operand int) int
	getComboOperand = func(operand int) int {
		switch operand {
		case 4:
			return registerA
		case 5:
			return registerB
		case 6:
			return registerC
		default:
			return operand
		}
	}

	output := ""
	expected := strings.Join(nums, ",")
	// a := 0b101_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000
	a := 0545311

	fmt.Println(a)

	for output != expected {
		// if a%1000000 == 0 {
		// 	fmt.Println(a)
		// 	fmt.Println(output)
		// }

		registerA = a
		registerB = 0
		registerC = 0

		// fmt.Println(registerA)

		outputs := make([]string, 0)

		instructionPointer := 0

		for instructionPointer < len(program) {
			fmt.Printf("registerA: %b\n", registerA)
			fmt.Printf("registerB: %b\n", registerB)
			fmt.Printf("registerC: %b\n", registerC)
			fmt.Println("----------------")

			opcode := program[instructionPointer]
			operand := program[instructionPointer+1]

			switch opcode {
			case 0: // adv division
				numerator := registerA
				denominator := math.Pow(2, float64(getComboOperand(operand)))
				registerA = int(math.Trunc(float64(numerator) / denominator))
				instructionPointer += 2
				continue
			case 1: // bxl bitwise XOR
				registerB = registerB ^ operand
				instructionPointer += 2
				continue
			case 2: // bst operand mod 8
				registerB = getComboOperand(operand) % 8
				instructionPointer += 2
				continue
			case 3: // jnz
				if registerA == 0 {
					instructionPointer += 2
					continue
				}
				instructionPointer = operand
				continue
			case 4: // bxc
				registerB = registerB ^ registerC
				instructionPointer += 2
				continue
			case 5: // out
				output := strconv.Itoa(getComboOperand(operand) % 8)
				outputs = append(outputs, output)
				instructionPointer += 2
				continue
			case 6: // bdv
				numerator := registerA
				denominator := math.Pow(2, float64(getComboOperand(operand)))
				registerB = int(math.Trunc(float64(numerator) / denominator))
				instructionPointer += 2
				continue
			case 7: // bdv
				numerator := registerA
				denominator := math.Pow(2, float64(getComboOperand(operand)))
				registerC = int(math.Trunc(float64(numerator) / denominator))
				instructionPointer += 2
				continue
			default:
				panic("unreachable")
			}
		}

		output = strings.Join(outputs, ",")
		fmt.Println("o: " + output)
		fmt.Println("e: " + expected)
		fmt.Println("====================================")
		a += 1

		break
	}

	fmt.Printf("registerA: %b\n", registerA)
	fmt.Printf("registerB: %b\n", registerB)
	fmt.Printf("registerC: %b\n", registerC)
	fmt.Println("o: " + output)
	fmt.Println("e: " + expected)

	return a - 1
}
