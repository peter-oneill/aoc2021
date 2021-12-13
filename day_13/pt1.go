// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type dot struct {
	x, y int
}

type fold struct {
	dir byte
	val int
}

func main() {
	input_sections := strings.Split(os.Args[1], "\n\n")

	var grid = build_initial_grid(input_sections[0])

	folds := build_fold_slice(input_sections[1])

	for _, f := range folds {
		grid = fold_grid(grid, f)
		break
	}
	print_grid(grid)

	num_dots := count_dots(grid)
	fmt.Println("num dots: ", num_dots)

	return
}

func build_initial_grid(dot_input_lines string) [][]bool {
	dot_strs := strings.Split(dot_input_lines, "\n")

	// Convert the input string to a list of dots, finding out the grid size along the way
	var dots []dot
	var max_x, max_y int = 0, 0

	for _, str := range dot_strs {
		coords := strings.Split(str, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		max_x = max(max_x, x)
		max_y = max(max_y, y)

		dots = append(dots, dot{x, y})
	}

	// Build a blank grid of the correct size
	var grid [][]bool

	for y := 0; y <= max_y; y++ {
		grid = append(grid, make([]bool, max_x+1))
	}

	// Fill the dots from the slice into the grid
	for _, dot := range dots {
		grid[dot.y][dot.x] = true
	}

	return grid
}

func build_fold_slice(fold_input_line string) []fold {
	fold_strs := strings.Split(fold_input_line, "\n")
	fold_regex := regexp.MustCompile("[xy]=\\d+$")

	var folds []fold

	for _, str := range fold_strs {
		matches := fold_regex.FindString(str)
		x_or_y := matches[0]
		val, _ := strconv.Atoi(matches[2:])
		folds = append(folds, fold{x_or_y, val})
	}

	return folds
}

func fold_grid(grid [][]bool, f fold) [][]bool {
	if f.dir == 'x' {
		for y, line := range grid {
			for ix_after_fold, val := range line[f.val+1:] {
				left_x := f.val - 1 - ix_after_fold
				grid[y][left_x] = grid[y][left_x] || val
			}
			grid[y] = grid[y][:f.val]
		}
	} else {
		for ix_after_fold, line := range grid[f.val+1:] {
			top_y := f.val - 1 - ix_after_fold
			for x, val := range line {
				grid[top_y][x] = grid[top_y][x] || val
			}
		}
		grid = grid[:f.val]
	}
	return grid
}

func count_dots(grid [][]bool) (num int) {
	num = 0
	for _, line := range grid {
		for _, d := range line {
			if d {
				num++
			}
		}
	}
	return
}

func print_grid(grid [][]bool) {
	for _, line := range grid {
		var line_str string
		for _, dot := range line {
			if dot {
				line_str = line_str + "#"
			} else {
				line_str = line_str + "."
			}
		}
		fmt.Println(string(line_str))
	}
	fmt.Println()
}

func max(a int, b int) (max int) {
	if a > b {
		return a
	}
	return b
}
