package day_02

import (
	"strings"
)

func parse(s string) [][]string {
	var out [][]string

	lines := strings.Split(s, "\n")

	for _, l := range lines {
		if l != "" {
			out = append(out, strings.Fields(l))
		}
	}

	return out
}
