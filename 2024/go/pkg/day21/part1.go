package day21

import (
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type cacheEntry struct {
	path  string
	depth int
}

type cache = map[cacheEntry]int

func createNumberPad() map[string]map[string][]string {
	numberPad := make(map[string]map[string][]string)

	numberPad["A"] = make(map[string][]string)
	numberPad["A"]["A"] = []string{""}
	numberPad["A"]["0"] = []string{"<"}
	numberPad["A"]["1"] = []string{"^<<", "<^<"}
	numberPad["A"]["2"] = []string{"^<,<^"}
	numberPad["A"]["3"] = []string{"^"}
	numberPad["A"]["4"] = []string{"^^<<", "^<<^", "<^^<"}
	numberPad["A"]["5"] = []string{"^^<", "<^^"}
	numberPad["A"]["6"] = []string{"^^"}
	numberPad["A"]["7"] = []string{"^^^<<", "<^^^<"}
	numberPad["A"]["8"] = []string{"^^^<", "<^^^"}
	numberPad["A"]["9"] = []string{"^^^"}

	numberPad["0"] = make(map[string][]string)
	numberPad["0"]["A"] = []string{">"}
	numberPad["0"]["0"] = []string{""}
	numberPad["0"]["1"] = []string{"^<"}
	numberPad["0"]["2"] = []string{"^"}
	numberPad["0"]["3"] = []string{"^>", ">^"}
	numberPad["0"]["4"] = []string{"^^<"}
	numberPad["0"]["5"] = []string{"^^"}
	numberPad["0"]["6"] = []string{"^^>", ">^^"}
	numberPad["0"]["7"] = []string{"^^^<", "^<^^", "^^<^"}
	numberPad["0"]["8"] = []string{"^^^"}
	numberPad["0"]["9"] = []string{"^^^>", ">^^^", "^>^^", "^^>^"}

	numberPad["1"] = make(map[string][]string)
	numberPad["1"]["A"] = []string{">>v"}
	numberPad["1"]["0"] = []string{">v"}
	numberPad["1"]["1"] = []string{""}
	numberPad["1"]["2"] = []string{">"}
	numberPad["1"]["3"] = []string{">>"}
	numberPad["1"]["4"] = []string{"^"}
	numberPad["1"]["5"] = []string{">^", "^>"}
	numberPad["1"]["6"] = []string{">>^", "^>>"}
	numberPad["1"]["7"] = []string{"^^"}
	numberPad["1"]["8"] = []string{"^^>", ">^^"}
	numberPad["1"]["9"] = []string{"^^>>", ">>^^", "^>>^"}

	numberPad["2"] = make(map[string][]string)
	numberPad["2"]["A"] = []string{">v", "v>"}
	numberPad["2"]["0"] = []string{"v"}
	numberPad["2"]["1"] = []string{"<"}
	numberPad["2"]["2"] = []string{""}
	numberPad["2"]["3"] = []string{">"}
	numberPad["2"]["4"] = []string{"^<", "<^"}
	numberPad["2"]["5"] = []string{"^"}
	numberPad["2"]["6"] = []string{"^>", ">^"}
	numberPad["2"]["7"] = []string{"^^<", "<^^"}
	numberPad["2"]["8"] = []string{"^^"}
	numberPad["2"]["9"] = []string{"^^>", ">^^"}

	numberPad["3"] = make(map[string][]string)
	numberPad["3"]["A"] = []string{"v"}
	numberPad["3"]["0"] = []string{"<v", "v<"}
	numberPad["3"]["1"] = []string{"<<"}
	numberPad["3"]["2"] = []string{"<"}
	numberPad["3"]["3"] = []string{""}
	numberPad["3"]["4"] = []string{"<<^", "^<<"}
	numberPad["3"]["5"] = []string{"^<", "<^"}
	numberPad["3"]["6"] = []string{"^"}
	numberPad["3"]["7"] = []string{"^^<<", "<<^^", "^<<^", "<^^<"}
	numberPad["3"]["8"] = []string{"^^<", "<^^"}
	numberPad["3"]["9"] = []string{"^^"}

	numberPad["4"] = make(map[string][]string)
	numberPad["4"]["A"] = []string{">>vv", "vv>>", ">vv>"}
	numberPad["4"]["0"] = []string{">vv"}
	numberPad["4"]["1"] = []string{"v"}
	numberPad["4"]["2"] = []string{">v", "v>"}
	numberPad["4"]["3"] = []string{">>v", "v>>"}
	numberPad["4"]["4"] = []string{""}
	numberPad["4"]["5"] = []string{">"}
	numberPad["4"]["6"] = []string{">>"}
	numberPad["4"]["7"] = []string{"^"}
	numberPad["4"]["8"] = []string{"^>", ">^"}
	numberPad["4"]["9"] = []string{"^>>", ">>^"}

	numberPad["5"] = make(map[string][]string)
	numberPad["5"]["A"] = []string{">vv", "vv>"}
	numberPad["5"]["0"] = []string{"vv"}
	numberPad["5"]["1"] = []string{"v<", "<v"}
	numberPad["5"]["2"] = []string{"v"}
	numberPad["5"]["3"] = []string{"v>"}
	numberPad["5"]["4"] = []string{"<"}
	numberPad["5"]["5"] = []string{""}
	numberPad["5"]["6"] = []string{">"}
	numberPad["5"]["7"] = []string{"^<", "<^"}
	numberPad["5"]["8"] = []string{"^"}
	numberPad["5"]["9"] = []string{"^>", ">^"}

	numberPad["6"] = make(map[string][]string)
	numberPad["6"]["A"] = []string{"vv"}
	numberPad["6"]["0"] = []string{"vv<", "<vv"}
	numberPad["6"]["1"] = []string{"v<<", "<<v"}
	numberPad["6"]["2"] = []string{"<v", "v<"}
	numberPad["6"]["3"] = []string{"v"}
	numberPad["6"]["4"] = []string{"<<"}
	numberPad["6"]["5"] = []string{"<"}
	numberPad["6"]["6"] = []string{""}
	numberPad["6"]["7"] = []string{"^<<", "<<^"}
	numberPad["6"]["8"] = []string{"^<", "<^"}
	numberPad["6"]["9"] = []string{"^"}

	numberPad["7"] = make(map[string][]string)
	numberPad["7"]["A"] = []string{">>vvv"}
	numberPad["7"]["0"] = []string{">vvv", "v>vv", "vv>v"}
	numberPad["7"]["1"] = []string{"vv"}
	numberPad["7"]["2"] = []string{"vv>", ">vv"}
	numberPad["7"]["3"] = []string{"vv>>", ">>vv"}
	numberPad["7"]["4"] = []string{"v"}
	numberPad["7"]["5"] = []string{">v", "v>"}
	numberPad["7"]["6"] = []string{"v>>", ">>v"}
	numberPad["7"]["7"] = []string{""}
	numberPad["7"]["8"] = []string{">"}
	numberPad["7"]["9"] = []string{">>"}

	numberPad["8"] = make(map[string][]string)
	numberPad["8"]["A"] = []string{"vvv>", ">vvv"}
	numberPad["8"]["0"] = []string{"vvv"}
	numberPad["8"]["1"] = []string{"vv<", "<vv"}
	numberPad["8"]["2"] = []string{"vv"}
	numberPad["8"]["3"] = []string{"vv>", ">vv"}
	numberPad["8"]["4"] = []string{"v<", "<v"}
	numberPad["8"]["5"] = []string{"v"}
	numberPad["8"]["6"] = []string{"v>", ">v"}
	numberPad["8"]["7"] = []string{"<"}
	numberPad["8"]["8"] = []string{""}
	numberPad["8"]["9"] = []string{">"}

	numberPad["9"] = make(map[string][]string)
	numberPad["9"]["A"] = []string{"vvv"}
	numberPad["9"]["0"] = []string{"vvv<", "<vvv"}
	numberPad["9"]["1"] = []string{"<<vv", "vv<<"}
	numberPad["9"]["2"] = []string{"<vv", "vv<"}
	numberPad["9"]["3"] = []string{"vv"}
	numberPad["9"]["4"] = []string{"<<v", "v<<"}
	numberPad["9"]["5"] = []string{"<v", "v<"}
	numberPad["9"]["6"] = []string{"v"}
	numberPad["9"]["7"] = []string{"<<"}
	numberPad["9"]["8"] = []string{"<"}
	numberPad["9"]["9"] = []string{""}

	return numberPad
}

func createDirectionalPad() map[string]map[string][]string {
	directionalPad := make(map[string]map[string][]string)

	directionalPad["A"] = make(map[string][]string)
	directionalPad["A"]["A"] = []string{""}
	directionalPad["A"]["^"] = []string{"<"}
	directionalPad["A"]["<"] = []string{"v<<", "<v<"}
	directionalPad["A"]["v"] = []string{"<v", "v<"}
	directionalPad["A"][">"] = []string{"v"}

	directionalPad["^"] = make(map[string][]string)
	directionalPad["^"]["A"] = []string{">"}
	directionalPad["^"]["^"] = []string{""}
	directionalPad["^"]["<"] = []string{"v<"}
	directionalPad["^"]["v"] = []string{"v"}
	directionalPad["^"][">"] = []string{"v>", ">v"}

	directionalPad["<"] = make(map[string][]string)
	directionalPad["<"]["A"] = []string{">>^", ">^>"}
	directionalPad["<"]["^"] = []string{">^"}
	directionalPad["<"]["<"] = []string{""}
	directionalPad["<"]["v"] = []string{">"}
	directionalPad["<"][">"] = []string{">>"}

	directionalPad["v"] = make(map[string][]string)
	directionalPad["v"]["A"] = []string{"^>", ">^"}
	directionalPad["v"]["^"] = []string{"^"}
	directionalPad["v"]["<"] = []string{"<"}
	directionalPad["v"]["v"] = []string{""}
	directionalPad["v"][">"] = []string{">"}

	directionalPad[">"] = make(map[string][]string)
	directionalPad[">"]["A"] = []string{"^", "^>"}
	directionalPad[">"]["^"] = []string{"<^", "^<"}
	directionalPad[">"]["<"] = []string{"<<"}
	directionalPad[">"]["v"] = []string{"<"}
	directionalPad[">"][">"] = []string{""}

	return directionalPad
}

func buildSequence(keys string, index int, path string, pad map[string]map[string][]string, result map[string]struct{}) {
	if index == len(keys) {
		result[path] = struct{}{}
		return
	}

	previous := "A"

	if index != 0 {
		previous = string(keys[index-1])
	}

	for _, v := range pad[previous][string(keys[index])] {
		buildSequence(keys, index+1, path+v+"A", pad, result)
	}
}

func shortestSequence(keys string, depth int, cache cache, pad map[string]map[string][]string) int {
	if depth == 0 {
		return len(keys)
	}

	cacheEntry := cacheEntry{path: keys, depth: depth}

	if length, ok := cache[cacheEntry]; ok {
		return length
	}

	subKeys := strings.Split(keys, "A")
	subKeys = subKeys[:len(subKeys)-1]

	total := 0

	for _, subKey := range subKeys {
		sequences := make(map[string]struct{})
		buildSequence(subKey+"A", 0, "", pad, sequences)

		minimum := math.MaxInt

		for sequence := range sequences {
			l := shortestSequence(sequence, depth-1, cache, pad)
			minimum = min(minimum, l)
		}

		total += minimum
	}

	cache[cacheEntry] = total

	return cache[cacheEntry]
}

func solve(input []string, depth int) int {
	numberPad := createNumberPad()
	directionalPad := createDirectionalPad()

	result := 0

	for _, code := range input {
		numericCode, _ := strconv.Atoi(code[0:3])
		numberPadSequences := make(map[string]struct{})
		buildSequence(code, 0, "", numberPad, numberPadSequences)

		cache := make(cache)
		minimum := math.MaxInt

		for numberPadSequence := range numberPadSequences {
			l := shortestSequence(numberPadSequence, depth, cache, directionalPad)
			minimum = min(minimum, l)
		}

		result += (minimum * numericCode)
	}

	return result
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)

	result := solve(input, 2)

	return result
}
