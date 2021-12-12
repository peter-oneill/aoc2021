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
	// input_lines := strings.Split(os.Args[1], "\n")
	input_lines := os.Args[1]
	var nodes []*node

	// for _, input_line := range input_lines {
	matches := regexp.MustCompile("([a-zA-Z]+)-([a-zA-Z]+)").FindAllStringSubmatch(input_lines, -1)
	for _, match := range matches {
		print_nodes(nodes)
		node1_name := match[1]
		node2_name := match[2]
		var node1_p *node
		var node2_p *node

		// fmt.Println("nodes", node1_name, " and", node2_name)

		// add each other's nodes to neighbours list
		for ii, n := range nodes {
			if n.name == node1_name {
				// fmt.Println("Found node 1", node1_name)
				node1_p = nodes[ii]
				// fmt.Println(node1_p)
				// fmt.Println(&node1_p)
			} else if n.name == node2_name {
				// fmt.Println("Found node 2", node2_name)
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
			// node1_p = &node1
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

		// fmt.Println(node1_p, node2_p)
		// fmt.Println(node1_p)
		// fmt.Println(&node1_p)
		// fmt.Println(node2_p)
		// fmt.Println(&node2_p)

		node1_p.neighbours = append(node1_p.neighbours, node2_p)
		node2_p.neighbours = append(node2_p.neighbours, node1_p)
		// fmt.Println(node1_p)
		// fmt.Println(&node1_p)
		// fmt.Println(node2_p)
		// fmt.Println(&node2_p)
		print_nodes(nodes)
		// fmt.Println(node1_p, node2_p)
	}

	// print_nodes(nodes)

	var path []*node
	// var start_node *node
	var start_node_ix int

	for node_ix, n := range nodes {
		if n.name == "start" {
			// start_node = &nodes[node_ix]
			start_node_ix = node_ix
			break
		}
	}

	// fmt.Println(start_node)
	// fmt.Printf("afer passing %p\n", start_node)

	print_nodes(nodes)
	// print_n
	visit_node(nodes, start_node_ix, path)

	// for ii := range start_node.neighbours {
	//     if start_node.neighbours[ii].visits_allowed != 0 {
	// 		path = visit_node(start_node.neighbours[ii], path)
	// 	}

	// }

	return
}

func visit_node(nodes []*node, index int, path []*node) {
	fmt.Printf("visiting %s at %p\n", (nodes)[index].name, (nodes)[index])
	print_node(nodes, index)
	switch nodes[index].visits_allowed {
	case 0:
		return
	case 1:
		nodes[index].visits_allowed = 0
	}
	path = append(path, nodes[index])
	if nodes[index].name == "end" {
		print_path(path)
	} else {
		for ii := range nodes[index].neighbours {
			var next_ix int
			for ix, n := range nodes {
				if n.name == nodes[index].neighbours[ii].name {
					next_ix = ix
				}
			}
			fmt.Println("next to visit: ", nodes[next_ix].name, nodes[next_ix])
			visit_node(nodes, next_ix, path)
		}
	}

	if nodes[index].visits_allowed == 0 {
		nodes[index].visits_allowed = 1
	}

	return
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
