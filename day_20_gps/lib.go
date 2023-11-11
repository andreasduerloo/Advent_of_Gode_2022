package day_20

import (
	"strconv"
	"strings"
)

// Returns a slice of integers and the length of that slice
func parse(s string) ([]int, int) {
	out := make([]int, 0)
	lines := strings.Split(s, "\n")

	for _, num := range lines {
		val, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		out = append(out, val)
	}
	return out, len(out)
}

func mix(lst []int) []int {
	out := copy(lst)
}

func swap(lst []int, index, val int) {
	if val > 0 { // We're moving to the right
		temp := lst[index]
		lst[index] = lst[(index+1)%len(lst)]
		lst[(index+1)%len(lst)]
	} else if val < 0 { // We're moving to the left
		//
	} else {
		return
	}
}
