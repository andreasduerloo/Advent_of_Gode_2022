package day_11

import (
	"advent2022/helpers"
	"strconv"
	"strings"
)

type monkey struct {
	items   []int
	action  rule
	div     int
	targets []int
	ifTrue  *monkey
	ifFalse *monkey
	thrown  int
}

type rule struct {
	operation string
	value     string
}

func parse(descs []string) []*monkey {
	monkeys := make([]*monkey, 0)

	for _, d := range descs {
		monkeys = append(monkeys, newMonkey(d))
	}

	// Populate the pointers
	for _, m := range monkeys {
		m.ifTrue = monkeys[m.targets[0]]
		m.ifFalse = monkeys[m.targets[1]]
	}

	return monkeys
}

func newMonkey(desc string) *monkey {
	lines := strings.Split(desc, "\n")
	ruleLine := strings.Fields(lines[2])

	trueStr := strings.Fields(lines[4])[len(strings.Fields(lines[4]))-1]
	falseStr := strings.Fields(lines[5])[len(strings.Fields(lines[5]))-1]

	trueInt, _ := strconv.Atoi(trueStr)
	falseInt, _ := strconv.Atoi(falseStr)

	out := monkey{
		items: helpers.ReGetInts(lines[1]),
		action: rule{
			operation: ruleLine[len(ruleLine)-2],
			value:     ruleLine[len(ruleLine)-1],
		},
		div:     helpers.ReGetInts(lines[3])[0],
		targets: []int{trueInt, falseInt},
	}

	return &out
}

func inspectAndThrow() {
	//
}

func applyRule() {
	//
}
