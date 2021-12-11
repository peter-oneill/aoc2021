// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input_lines := strings.Split(os.Args[1], "\n")
	var octopus_grid = [][]octopus{}

	for y, line := range input_lines {
		octopus_grid = append(octopus_grid, []octopus{})
		for _, digit := range line {
			value, _ := strconv.Atoi(string(digit))
			octopus_grid[y] = append(octopus_grid[y], octopus{value, false})
		}
	}

	num_steps := 100
	var flashes = 0
	for step := 0; step < num_steps; step++ {
		// Increment every octopus' value
		for y, octopus_line := range octopus_grid {
			for x := range octopus_line {
				increase_energy(octopus_grid, x, y)
			}
		}

		for y, octopus_line := range octopus_grid {
			for x, octo := range octopus_line {
				if octo.flashed {
					flashes++
					octopus_grid[y][x].energy = 0
					octopus_grid[y][x].flashed = false

				} else if octo.energy > 9 {
					fmt.Println("err!")
				}
			}
		}
	}

	fmt.Println(flashes)

	return
}

func print_grid(grid [][]octopus) {
	for _, octo_line := range grid {
		for _, octo := range octo_line {
			fmt.Printf("%d", octo.energy)
		}
		fmt.Printf("\n")
	}
	fmt.Println("")
}

type octopus struct {
	energy  int
	flashed bool
}

func increase_energy(octopus_grid [][]octopus, x int, y int) {
	octopus_grid[y][x].energy++

	if !octopus_grid[y][x].flashed && octopus_grid[y][x].energy > 9 {
		octopus_grid[y][x].flashed = true

		if x > 0 {
			if y > 0 {
				increase_energy(octopus_grid, x-1, y-1)
			}
			increase_energy(octopus_grid, x-1, y)
			if y < len(octopus_grid)-1 {
				increase_energy(octopus_grid, x-1, y+1)
			}
		}

		if y > 0 {
			increase_energy(octopus_grid, x, y-1)
		}
		if y < len(octopus_grid)-1 {
			increase_energy(octopus_grid, x, y+1)
		}

		if x < len(octopus_grid)-1 {
			if y > 0 {
				increase_energy(octopus_grid, x+1, y-1)
			}
			increase_energy(octopus_grid, x+1, y)
			if y < len(octopus_grid)-1 {
				increase_energy(octopus_grid, x+1, y+1)
			}
		}
	}
}
