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
	visits_allowed int // -1: any number, 1: lower case, 0: lower case after visit
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

			var visits_allowed = -1
			if strings.ToLower(node1_name) == node1_name {
				visits_allowed = 1
			}
			node1 := node{node1_name, nil, visits_allowed}
			nodes = append(nodes, &node1)
			node1_p = nodes[len(nodes)-1]
		}
		if node2_p == nil {
			var visits_allowed = -1
			if strings.ToLower(node2_name) == node2_name {
				visits_allowed = 1
			}
			node2 := node{node2_name, nil, visits_allowed}
			nodes = append(nodes, &node2)
			node2_p = nodes[len(nodes)-1]
		}

		node1_p.neighbours = append(node1_p.neighbours, node2_p)
		node2_p.neighbours = append(node2_p.neighbours, node1_p)
	}

	var path []*node
	var start_node_ix int

	for node_ix, n := range nodes {
		if n.name == "start" {
			start_node_ix = node_ix
			break
		}
	}

	paths_found := visit_node(nodes, start_node_ix, path, 0)
	fmt.Println("Paths found: ", paths_found)

	return
}

func visit_node(nodes []*node, index int, path []*node, paths_found int) int {
	var new_paths_found int = paths_found
	switch nodes[index].visits_allowed {
	case 0:
		return paths_found
	case 1:
		nodes[index].visits_allowed = 0
	}
	path = append(path, nodes[index])
	if nodes[index].name == "end" {
		// print_path(path)
		new_paths_found++
	} else {
		for ii := range nodes[index].neighbours {
			var next_ix int
			for ix, n := range nodes {
				if n.name == nodes[index].neighbours[ii].name {
					next_ix = ix
				}
			}
			new_paths_found = visit_node(nodes, next_ix, path, new_paths_found)
		}
	}

	if nodes[index].visits_allowed == 0 {
		nodes[index].visits_allowed = 1
	}

	return new_paths_found
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
