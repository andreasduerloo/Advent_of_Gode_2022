package day_21

import (
	"strconv"
	"strings"
)

type monkey struct {
	value     int
	cstr      [2]string  // The 'names' of the children
	children  [2]*monkey // Pointers to the children
	operation string     // +, -, *, or /

	humnbranch bool    // Is this the branch that has humn in it?
	parent     *monkey // Pointer to the parent
}

func parse(s string) map[string]*monkey {
	lines := strings.Split(string(s), "\n")
	monkeyMap := make(map[string]*monkey)

	for _, line := range lines {
		monkify(line, &monkeyMap)
	}

	for _, v := range monkeyMap { // Populate the pointers
		if v.value == 0 { // Check not needed for first star
			v.children[0] = monkeyMap[v.cstr[0]]
			v.children[1] = monkeyMap[v.cstr[1]]

			// Second star
			monkeyMap[v.cstr[0]].parent = v
			monkeyMap[v.cstr[1]].parent = v
		}
	}

	// Needed for the second star
	m := monkeyMap["humn"]
	m.humnbranch = true
	for !monkeyMap["root"].humnbranch {
		m = m.parent
		m.humnbranch = true
	}

	return monkeyMap
}

func monkify(l string, m *map[string]*monkey) { // Do we have to pass a pointer to the map?
	if l != "" {
		elems := strings.Split(l, " ")

		switch len(elems) {
		case 2:
			val, _ := strconv.Atoi(elems[1])
			newm := monkey{
				value: val,
			}

			(*m)[elems[0][:len(elems[0])-1]] = &newm
		case 4:
			newm := monkey{
				cstr:      [2]string{elems[1], elems[3]},
				operation: elems[2],
			}

			(*m)[elems[0][:len(elems[0])-1]] = &newm
		}
	}
}

func solveTree(m *monkey) int { // RECURSION! Parallelize?
	if m.value != 0 {
		return m.value
	} else {
		switch m.operation {
		case "+":
			m.value = solveTree(m.children[0]) + solveTree(m.children[1])
		case "-":
			m.value = solveTree(m.children[0]) - solveTree(m.children[1])
		case "*":
			m.value = solveTree(m.children[0]) * solveTree(m.children[1])
		case "/":
			m.value = solveTree(m.children[0]) / solveTree(m.children[1])
		}

		return m.value
	}
}

func makeEqual(m *monkey) int { // We have to go down, at each step we know what the value of m should be
	if m.parent == nil { // This is "root". Set the value of the "humn" branch to that of the other branch
		if m.children[0].humnbranch { // The left child is an ancestor of "humn"
			m.children[0].value = m.children[1].value
			return makeEqual(m.children[0])
		} else { // The right one is
			m.children[0].value = m.children[1].value
			return makeEqual(m.children[1])
		}

	} else if m.children[0] == nil { // This is "humn"
		return m.value

	} else { // Work out what the value of the "humn" child should be, so that this monkey's value is correct
		var ancestor *monkey
		var othermonkey *monkey
		if m.children[0].humnbranch { // The left child is an ancestor of "humn"
			ancestor = m.children[0]
			othermonkey = m.children[1]
		} else { // The right one is
			ancestor = m.children[1]
			othermonkey = m.children[0]
		}

		switch m.operation {
		case "+":
			ancestor.value = m.value - othermonkey.value
		case "-": // Order matters
			if m.value > othermonkey.value {
				ancestor.value = m.value + othermonkey.value
			} else {
				ancestor.value = othermonkey.value - m.value
			}
		case "*":
			ancestor.value = m.value / othermonkey.value
		case "/": // Order matters
			if m.value > othermonkey.value {
				ancestor.value = m.value * othermonkey.value
			} else {
				ancestor.value = othermonkey.value / m.value
			}
		}

		return makeEqual(ancestor)
	}
}
