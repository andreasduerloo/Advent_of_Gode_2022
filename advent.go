package main

import (
	day_01 "advent2022/day_01_calories"
	day_02 "advent2022/day_02_rockpaperscissors"
	day_04 "advent2022/day_04_cleanup"
	day_11 "advent2022/day_11_monkeys"
	day_12 "advent2022/day_12_hillclimbing"
	day_13 "advent2022/day_13_packets"
	day_21 "advent2022/day_21_monkeymath"
	day_25 "advent2022/day_25_snafu"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No argument was passed - exiting.")
		return
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The argument is not an integer - exiting.")
	}

	fmt.Println("Solutions for day", day)

	switch day {
	case 1:
		first, second := day_01.Solve()
		fmt.Println(first, second)
	case 2:
		first, second := day_02.Solve()
		fmt.Println(first, second)
	case 4:
		first, second := day_04.Solve()
		fmt.Println(first, second)
	case 11:
		first, second := day_11.Solve()
		fmt.Println(first, second)
	case 12:
		first, second := day_12.Solve()
		fmt.Println(first, second)
	case 13:
		first, second := day_13.Solve()
		fmt.Println(first, second)
	// case 20:
	// first, second := day_20.Solve()
	// fmt.Println(first, second)
	case 21:
		first, second := day_21.Solve()
		fmt.Println(first, second)
	case 25:
		first, second := day_25.Solve()
		fmt.Println(first, second)
	default:
		fmt.Println("That's either not a valid day, or it has not been solved (yet!)")
	}

}
