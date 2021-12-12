// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type node struct {
	name           string
	neighbours     []*node
	visits_allowed int // -1: any number, 2: lower case
}

func main() {
	input_lines := os.Args[1]
	var nodes []*node

	matches := regexp.MustCompile("([a-zA-Z]+)-([a-zA-Z]+)").FindAllStringSubmatch(input_lines, -1)
	for _, match := range matches {
		node1_name := match[1]
		node2_name := match[2]
		var node1_p *node
		var node2_p *node

		for ii, n := range nodes {
			if n.name == node1_name {
				node1_p = nodes[ii]
			} else if n.name == node2_name {
				node2_p = nodes[ii]
			}
		}

		if node1_p == nil {
			node1 := new_node(node1_name)
			node1_p = &node1
			nodes = append(nodes, node1_p)
		}
		if node2_p == nil {
			node1 := new_node(node2_name)
			node2_p = &node1
			nodes = append(nodes, node2_p)
		}

		node1_p.neighbours = append(node1_p.neighbours, node2_p)
		node2_p.neighbours = append(node2_p.neighbours, node1_p)
	}

	var path []*node
	var start_node *node

	for _, n := range nodes {
		if n.name == "start" {
			start_node = n
			break
		}
	}

	paths_found := visit_node(start_node, path, 0, false)
	fmt.Println("Paths found: ", paths_found)

	return
}

func visit_node(n *node, path []*node, paths_found int, had_double_visit bool) int {
	if n.visits_allowed == 0 {
		return paths_found
	}

	if n.name == "end" {
		print_path(append(path, n))
		return paths_found + 1
	}

	if had_double_visit && strings.ToLower(n.name) == n.name && n.visits_allowed == 1 {
		return paths_found
	}

	// We're allowed to visit this node!
	switch strings.ToLower(n.name) {
	case "start":
		n.visits_allowed = 0
	case n.name:
		switch n.visits_allowed {
		case 1:
			n.visits_allowed = 0
			had_double_visit = true
		case 2:
			n.visits_allowed = 1
		}
	}

	path = append(path, n)

	for _, n := range n.neighbours {
		paths_found = visit_node(n, path, paths_found, had_double_visit)
	}

	if strings.ToLower(n.name) == n.name {
		n.visits_allowed++
	}

	return paths_found
}

func new_node(node_name string) node {
	var visits_allowed = -1
	if strings.ToLower(node_name) == node_name {
		visits_allowed = 2
	}
	return node{node_name, nil, visits_allowed}
}

func print_path(path []*node) {
	fmt.Printf("New path: ")
	for _, n := range path {
		fmt.Printf("%s,", n.name)
	}
	fmt.Printf("\n")
}

func print_node(nodes []*node, ix int) {
	fmt.Printf("node %s at %p, visits: %d neighbours: ", nodes[ix].name, &nodes[ix], nodes[ix].visits_allowed)
	for _, neigh := range nodes[ix].neighbours {
		fmt.Printf("%s,", neigh.name)
	}
	fmt.Printf("\n")
}

func print_nodes(nodes []*node) {
	for ii := range nodes {
		print_node(nodes, ii)
	}
}
