package day_04

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/04.txt")
	if err != nil {
		fmt.Println("Error getting input for day 4 - exiting")
		return 0, 0
	}

	bounds := parse(string(input))
	var firstStar, secondStar int

	for _, pair := range bounds {
		fullyOverlaps(pair, &firstStar)
		overlaps(pair, &secondStar)
	}

	return firstStar, secondStar
}
