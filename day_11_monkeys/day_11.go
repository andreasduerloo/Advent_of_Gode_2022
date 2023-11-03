package day_11

import (
	"fmt"
	"os"
	"strings"
)

func Solve() (int, int) {
	input, err := os.ReadFile("inputs/12.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	monkeys := parse(strings.Split(string(input), "\n\n"))

	return 0, 0
}
