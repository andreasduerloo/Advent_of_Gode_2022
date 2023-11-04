package day_11

import (
	"advent2022/helpers"
	"slices"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	action    rule
	div       int
	targets   []int
	ifTrue    *monkey
	ifFalse   *monkey
	inspected int
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

// This function takes a pointer to a monkey and a boolean to indicate what star we're solving.
// True is the second star
func inspectAndThrowItems(m *monkey, s bool, magicNum int) {
	for _, item := range m.items {
		// Inspect the item: apply the rule
		worryVal := applyRule(m, item)
		if worryVal > magicNum { // Math magic
			worryVal = worryVal % magicNum
		}
		m.inspected += 1

		// Brief relief
		if !s {
			worryVal = worryVal / 3
		}

		// Test and throw
		if worryVal%m.div == 0 {
			m.ifTrue.items = append(m.ifTrue.items, worryVal)
		} else {
			m.ifFalse.items = append(m.ifFalse.items, worryVal)
		}
	}

	// Empty this monkey's item queue
	m.items = make([]int, 0)
}

func applyRule(m *monkey, v int) int {
	var otherVal int
	switch m.action.value {
	case "old":
		otherVal = v
	default:
		otherVal, _ = strconv.Atoi(m.action.value)
	}

	switch m.action.operation {
	case "+":
		v = v + otherVal
	case "*":
		v = v * otherVal
	}
	return v
}

func monkeyBusiness(monkeys []*monkey) int {
	// Order the slice of monkeys by 'inspected'
	slices.SortFunc(monkeys, func(a, b *monkey) int { return b.inspected - a.inspected })

	// Take the top two and multiply their values
	return monkeys[0].inspected * monkeys[1].inspected
}

// Returns the product of all monkey.div values
// Needed for moludar arithmetic alchemy
func magicNumber(monkeys []*monkey) int {
	out := 1

	for _, m := range monkeys {
		out *= m.div
	}

	return out
}
