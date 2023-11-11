package day_20

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/20.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	numList, numLen := parse(input)

	return 0, 0
}
