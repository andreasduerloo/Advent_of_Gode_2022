package day_13

import (
	"fmt"
	"os"
	"slices"
)

func Solve() (int, int) {
	input, err := os.ReadFile("inputs/13.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	packets := parse(string(input))
	var correct int

	for i := 0; i < len(packets); i += 2 {
		if comparePackets(packets[i], packets[i+1]) == Ok {
			correct += (i / 2) + 1
		}
	}

	packets = appendDividers(packets)

	slices.SortFunc(packets, orderPackets)

	dividers := findDividers(packets)

	return correct, dividers
}
