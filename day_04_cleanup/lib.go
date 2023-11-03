package day_04

import (
	"advent2022/helpers"
	"strings"
)

func parse(s string) [][]int {
	lines := strings.Split(s, "\n")
	bounds := make([][]int, 0)

	for _, line := range lines {
		bounds = append(bounds, helpers.ReGetInts(line))
	}

	return bounds
}

func fullyOverlaps(b []int, c *int) {
	if len(b) > 0 {
		if b[0] <= b[2] && b[1] >= b[3] || b[0] >= b[2] && b[1] <= b[3] {
			*c += 1
		}
	}
}

func overlaps(b []int, c *int) {
	if len(b) > 0 {
		if !(b[1] < b[2] || b[0] > b[3]) {
			*c += 1
		}
	}
}
