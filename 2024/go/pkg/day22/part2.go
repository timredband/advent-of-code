package day22

import (
	"math"
	"os"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type priceChange struct {
	price  int
	change int
}

type Range struct {
	first  int
	second int
	third  int
	fourth int
}

func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	rangesByTotal := make(map[Range]int)

	for i := range input {
		priceChanges := make([]priceChange, 2000)

		rangesByTotalForInput := make(map[Range]int)

		secret, _ := strconv.Atoi(input[i])
		price := secret % 10

		for i := range 2000 {
			secret = calculateNextSecret(secret)
			previousPrice := price
			price = secret % 10
			diff := price - previousPrice

			priceChanges[i] = priceChange{price: price, change: diff}
		}

		for i := 3; i < len(priceChanges); i += 1 {
			r := Range{
				first:  priceChanges[i-3].change,
				second: priceChanges[i-2].change,
				third:  priceChanges[i-1].change,
				fourth: priceChanges[i].change,
			}

			if _, ok := rangesByTotalForInput[r]; !ok {
				rangesByTotalForInput[r] += priceChanges[i].price
			}
		}

		for r, total := range rangesByTotalForInput {
			rangesByTotal[r] += total
		}
	}

	maximumPrice := math.MinInt

	for _, total := range rangesByTotal {
		maximumPrice = max(maximumPrice, total)
	}

	return maximumPrice
}
