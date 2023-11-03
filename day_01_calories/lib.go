package day_01

import (
	"strconv"
	"strings"
)

func parse(s string) []int {
	outElves := make([]int, 0)
	elves := strings.Split(s, "\n\n")

	for _, elf := range elves {
		calCount := 0
		items := strings.Split(elf, "\n")

		for _, valStr := range items {
			val, err := strconv.Atoi(valStr)
			if err != nil {
				continue
			}
			calCount += val
		}

		outElves = append(outElves, calCount)
	}

	return outElves
}
