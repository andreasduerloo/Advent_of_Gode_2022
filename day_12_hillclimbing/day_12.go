package day_12

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

	grid, start, end := buildGrid(strings.Split(string(input), "\n"))

	first := bfs(start, end, true)
	reset(grid)
	second := bfs(end, start, false)

	return first, second
}
