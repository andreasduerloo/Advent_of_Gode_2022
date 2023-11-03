package helpers

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// A queue can be implemented by using a slice and appending items to the end.
// Dequeue returns the first item in the queue and the rest of the queue (or an empty slice)
func Dequeue[T any](q []T) (T, []T) {
	out := q[0]

	if len(q) == 1 {
		return out, make([]T, 0)
	} else {
		return out, q[1:]
	}
}

// Completely replaced by function ReGetInts - see below
// Returns all the integers found in a string as a slice
// The integers need to be separated from any non-integer runes by whitespace
func GetAllInts(s string) []int {
	split := strings.Split(s, " ")
	out := make([]int, 0)

	for _, word := range split {
		val, err := strconv.Atoi(word)
		if err != nil {
			continue
		}
		out = append(out, val)
	}
	return out
}

// Returns all the integers found in a string as a slice
// Based on a RegEx, so the integers do not need to be separated from non-integer runes
func ReGetInts(s string) []int {
	re := regexp.MustCompile(`[0-9]+`)
	matches := re.FindAllString(s, -1)

	ints := make([]int, 0)

	for _, match := range matches {
		val, err := strconv.Atoi(match)
		if err != nil {
			continue
		}
		ints = append(ints, val)
	}

	return ints
}

type point struct {
	x int
	y int
}

// Returns the manhattan distance between two points
func ManhattanDistance(a, b point) int {
	return slices.Max([]int{a.x, b.x}) - slices.Min([]int{a.x, b.x}) + slices.Max([]int{a.y, b.y}) - slices.Min([]int{a.y, b.y})
}

// When using map[T]struct{} as a set, this is a somewhat nicer way to check for membership
// Might be overkill (if _, present := map[key]; present)
func Member[T comparable](k T, m map[T]struct{}) bool {
	_, present := m[k]
	return present
}

// Four ways to represent a grid
// - A single slice
// - A 2D array
// - A Map
// - A graph

func GridAsMap() {
	//
}

// Returns a single slice contianing all the positions in the grid, as well as the length of a row
// Using the row length, neighbors can be calculated (i - 1, i + 1, i - length, i + length)
// ints bool -> also do Atoi
func GridAsSlice() {
	//
}

func GridAs2DSlice() {
	//
}
