package day_14

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("inputs/14.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	state, lowest := parse(string(input))

	fmt.Println(lowest)
	fmt.Println(len(state))

	// First star
	abyss := false
	first := 0

	for abyss == false {
		abyss = drop(state, lowest)
		if !abyss {
			first += 1
		}
	}

	// Second star
	state, lowest = parse(string(input))
	second := 0

	for !state[point{x: 500, y: 0}] {
		drop2(state, lowest)
		second += 1
	}

	return first, second
}
