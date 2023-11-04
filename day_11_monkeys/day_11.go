package day_11

import (
	"fmt"
	"os"
	"strings"
)

func Solve() (int, int) {
	input, err := os.ReadFile("inputs/11.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	monkeys := parse(strings.Split(string(input), "\n\n"))
	magicNum := magicNumber(monkeys)

	// First star - twenty rounds
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			inspectAndThrowItems(m, false, magicNum)
		}
	}

	first := monkeyBusiness(monkeys)

	// Second star - reset the monkeys
	monkeys = parse(strings.Split(string(input), "\n\n"))

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			inspectAndThrowItems(m, true, magicNum)
		}
	}

	second := monkeyBusiness(monkeys)

	return first, second
}
