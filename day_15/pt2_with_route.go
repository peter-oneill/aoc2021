// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y                      int
	risk                      int
	cheapest_path             int
	cheapest_path_predecessor *position
	on_path                   bool
}

func main() {
	input := os.Args[1]

	var grid [][]position

	for y, line := range strings.Split(input, "\n") {
		grid = append(grid, make([]position, 0))
		for x, strval := range line {
			num_val, _ := strconv.Atoi(string(strval))
			new_pos := position{x, y, num_val, -1, nil, false}
			grid[y] = append(grid[y], new_pos)
		}
	}
	orig_x_len := len(grid[0])
	orig_y_len := len(grid)

	orig_grid := grid

	grid = append(grid, orig_grid...)
	grid = append(grid, orig_grid...)
	grid = append(grid, orig_grid...)
	grid = append(grid, orig_grid...)

	for y, line := range grid {
		grid[y] = append(grid[y], line...)
		grid[y] = append(grid[y], line...)
		grid[y] = append(grid[y], line...)
		grid[y] = append(grid[y], line...)

		for x := range grid[y] {
			node := &grid[y][x]
			node.risk = node.risk + x/orig_x_len + y/orig_y_len
			if node.risk > 9 {
				node.risk = node.risk - 9
			}
			node.x = x
			node.y = y
		}
	}

	// print_grid(grid)

	start_x, start_y := 0, 0
	end_x := len(grid[0]) - 1
	end_y := len(grid) - 1

	starting_node := &grid[start_y][start_x]
	starting_node.cheapest_path = 0

	var positions_to_hunt_from = []*position{&grid[start_y][start_x]}

	for {
		if len(positions_to_hunt_from) == 0 {
			break
		}
		var next_positions []*position
		for _, pos_ptr := range positions_to_hunt_from {
			x := pos_ptr.x
			y := pos_ptr.y
			if x > 0 {
				append_if_better_path(&next_positions, pos_ptr, &grid[pos_ptr.y][pos_ptr.x-1])
			}
			if x < end_x {
				append_if_better_path(&next_positions, pos_ptr, &grid[pos_ptr.y][pos_ptr.x+1])
			}
			if y > 0 {
				append_if_better_path(&next_positions, pos_ptr, &grid[pos_ptr.y-1][pos_ptr.x])
			}
			if y < end_y {
				append_if_better_path(&next_positions, pos_ptr, &grid[pos_ptr.y+1][pos_ptr.x])
			}
		}
		positions_to_hunt_from = next_positions
	}

	fmt.Println("cost to bottom right: ", grid[end_y][end_x].cheapest_path)

	print_route(grid)

	return
}

func append_if_better_path(list *[]*position, pos_ptr *position, neighbour *position) {
	this_route_risk := pos_ptr.cheapest_path + neighbour.risk
	if this_route_risk < neighbour.cheapest_path || neighbour.cheapest_path == -1 {
		neighbour.cheapest_path = this_route_risk
		neighbour.cheapest_path_predecessor = pos_ptr
		(*list) = append(*list, neighbour)
	}
}

func print_grid(grid [][]position) {
	for _, line := range grid {
		for _, x := range line {
			fmt.Printf("%d", x.risk)
		}
		fmt.Println()
	}
}

func print_route(grid [][]position) {
	// Start at the final node
	var working_node *position = &grid[len(grid)-1][len(grid[0])-1]
	for {
		working_node.on_path = true
		if working_node.x+working_node.y == 0 {
			break
		}
		working_node = working_node.cheapest_path_predecessor
	}

	print_grid_with_path(grid)
	return
}

func print_grid_with_path(grid [][]position) {
	for _, line := range grid {
		var buf bytes.Buffer
		for _, x := range line {
			var char string
			if x.on_path {
				char = "X"
			} else {
				char = fmt.Sprintf("%d", x.risk)
			}
			buf.WriteString(char)
		}

		fmt.Println(buf.String())
	}
}
