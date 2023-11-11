package day_14

import (
	"advent2022/helpers"
	"slices"
	"strings"
)

type point struct {
	x int
	y int
}

func parse(input string) (map[point]bool, int) {
	state := make(map[point]bool)
	var lowest int

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		ints := helpers.ReGetInts(line)

		for i := 0; i < len(ints)-3; i += 2 {
			if ints[i] == ints[i+2] { // Vertical
				for j := slices.Min([]int{ints[i+1], ints[i+3]}); j <= slices.Max([]int{ints[i+1], ints[i+3]}); j++ {
					state[point{x: ints[i], y: j}] = true
					if j > lowest {
						lowest = j
					}
				}
			} else { // Horizontal
				if ints[i+1] > lowest {
					lowest = ints[i+1]
				}
				for j := slices.Min([]int{ints[i], ints[i+2]}); j <= slices.Max([]int{ints[i], ints[i+2]}); j++ {
					state[point{x: j, y: ints[i+1]}] = true
				}
			}
		}
	}

	return state, lowest
}

// Returns true once a grain of sand starts falling into the abyss
func drop(state map[point]bool, lowest int) bool {
	// All grains of sand start at 500,0
	pos := point{x: 500, y: 0}
	stopped := false

	for !stopped {
		if pos.y > lowest {
			return true
		}

		if !state[point{x: pos.x, y: pos.y + 1}] {
			pos.y = pos.y + 1
		} else if !state[point{x: pos.x - 1, y: pos.y + 1}] {
			pos.x = pos.x - 1
			pos.y = pos.y + 1
		} else if !state[point{x: pos.x + 1, y: pos.y + 1}] {
			pos.x = pos.x + 1
			pos.y = pos.y + 1
		} else {
			state[pos] = true
			stopped = true
		}
	}

	return false
}

func drop2(state map[point]bool, lowest int) {
	// All grains of sand start at 500,0
	pos := point{x: 500, y: 0}
	stopped := false

	for !stopped {
		if !state[point{x: pos.x, y: pos.y + 1}] && pos.y+1 < 2+lowest {
			pos.y = pos.y + 1
		} else if !state[point{x: pos.x - 1, y: pos.y + 1}] && pos.y+1 < 2+lowest {
			pos.x = pos.x - 1
			pos.y = pos.y + 1
		} else if !state[point{x: pos.x + 1, y: pos.y + 1}] && pos.y+1 < 2+lowest {
			pos.x = pos.x + 1
			pos.y = pos.y + 1
		} else {
			state[pos] = true
			stopped = true
		}
	}
}
