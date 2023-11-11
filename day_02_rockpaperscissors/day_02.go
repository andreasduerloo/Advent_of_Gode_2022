package day_02

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/02.txt")
	if err != nil {
		fmt.Println("Error getting input - exiting")
		return 0, 0
	}

	instructions := parse(string(input))

	ruleSet := map[string]map[string][]int{
		"X": map[string][]int{"A": []int{4, 3}, "B": []int{1, 1}, "C": []int{7, 2}},
		"Y": map[string][]int{"A": []int{8, 4}, "B": []int{5, 5}, "C": []int{2, 6}},
		"Z": map[string][]int{"A": []int{3, 8}, "B": []int{9, 9}, "C": []int{6, 7}},
	}

	var first, second int

	for _, move := range instructions {
		first += ruleSet[move[1]][move[0]][0]
		second += ruleSet[move[1]][move[0]][1]
	}

	return first, second
}
