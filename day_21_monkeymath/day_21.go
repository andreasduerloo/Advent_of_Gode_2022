package day_21

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/21.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	monkeyMap := parse(string(input))

	return solveTree(monkeyMap["root"]), makeEqual(monkeyMap["root"])
}
