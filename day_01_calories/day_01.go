package day_01

import (
	"fmt"
	"os"
	"slices"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/01.txt")
	if err != nil {
		fmt.Println("Error getting input for day 1 - exiting")
		return 0, 0
	}

	elves := parse(string(input))
	slices.SortFunc(elves, func(x, y int) int { return y - x })

	return elves[0], (elves[0] + elves[1] + elves[2])
}
