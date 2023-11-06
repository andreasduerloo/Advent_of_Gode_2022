package day_25

import (
	"fmt"
	"os"
	"strings"
)

func Solve() (string, int) {
	input, err := os.ReadFile("./inputs/25.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return "N/A", 0
	}

	var sumDecimal int

	for _, s := range strings.Split(string(input), "\n") {
		sumDecimal += snafuToDecimal(s)
	}

	return decimalToSnafu(sumDecimal), 0
}
