package day_12

import (
	"advent2022/helpers"
)

type node struct {
	value     rune
	visited   bool
	neighbors []*node
	steps     int
}

func newNode(value rune) *node {
	out := node{
		value:     value,
		visited:   false,
		neighbors: make([]*node, 0),
		steps:     0,
	}
	return &out
}

// This function creates a graph of nodes linked to each other.
// It returns the entire graph as a slice containing all the nodes (we'll need this to reset the graph)
// as well as the 'start' node, and the end node.
func buildGrid(lines []string) ([]*node, *node, *node) {
	var graph []*node
	var start, end *node
	var width, height int

	for i, l := range lines {
		if l != "" {
			for j, r := range l {
				graph = append(graph, newNode(r))

				if r == 'S' {
					start = graph[len(graph)-1]
					start.value = 'a'
				} else if r == 'E' {
					end = graph[len(graph)-1]
					end.value = 'z'
				}

				if j > width {
					width = j
				}
			}
		}
		if i > height {
			height = i
		}
	}

	// Catch an off-by-one error
	width += 1

	// Populate the neighbor pointers
	for i, n := range graph {
		if i%width != 0 { // Check for left edge
			n.neighbors = append(n.neighbors, graph[i-1])
		}
		if i%width != width-1 && i != width*height { // Check for right edge
			n.neighbors = append(n.neighbors, graph[i+1])
		}
		if i >= width { // Check for top edge
			n.neighbors = append(n.neighbors, graph[i-width])
		}
		if i < (width*height)-width { // Check for bottom edge
			n.neighbors = append(n.neighbors, graph[i+width])
		}
	}

	return graph, start, end
}

// The meat and potatoes of the solution: breadth-first search.
// Takes a starting node and a destination node as input, as well as a boolean indicating wether we are climbing or descending
// True means climbing, false means descending
// Returns the shortest distance between the start and the end
func bfs(start *node, end *node, climb bool) int {
	var queue []*node
	current := start
	current.visited = true

	switch climb {
	case true:
		for end.steps == 0 { // In other words, as long as we have not visited the destination node yet
			for _, n := range current.neighbors {
				if n.value <= current.value+1 && !n.visited {
					n.visited = true
					n.steps = current.steps + 1
					queue = append(queue, n)
				}
			}

			if len(queue) > 0 {
				current, queue = helpers.Dequeue(queue)
			}
		}
		return end.steps
	case false:
		for current.value != 'a' { // Not quite low enough yet
			for _, n := range current.neighbors {
				if n.value >= current.value-1 && !n.visited {
					n.visited = true
					n.steps = current.steps + 1
					queue = append(queue, n)
				}
			}

			if len(queue) > 0 {
				current, queue = helpers.Dequeue(queue)
			}
		}
	}
	return current.steps
}

func reset(grid []*node) {
	for _, n := range grid {
		n.visited = false
		n.steps = 0
	}
}
