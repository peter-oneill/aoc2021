// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	// "fmt"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input_lines := strings.Split(os.Args[1], "\n")
	var heights [][]point

	for y, input_line := range input_lines {
		if y == len(heights) {
			heights = append(heights, []point{})
		}

		heights[y] = append(heights[y], point{-1, y, 10, -1})

		for x, input_digit := range []rune(input_line) {
			height, _ := strconv.Atoi(string(input_digit))
			heights[y] = append(heights[y], point{x, y, height, -1})
		}

		heights[y] = append(heights[y], point{len(heights[0]), y, 10, -1})
	}

	var top_border = make([]point, len(heights[0]))

	for x, _ := range top_border {
		top_border[x] = point{x, -1, 10, -1}
	}

	heights = append(heights, top_border)

	var bottom_border = make([]point, len(heights[0]))

	for x, _ := range bottom_border {
		bottom_border[x] = point{x, -1, 10, -1}
	}

	heights = append([][]point{bottom_border}, heights...)

	var basin_ix = 0
	var basins []basin

	for y, height_line := range heights[1 : len(heights)-1] {

		for x, point := range height_line[1 : len(height_line)-1] {
			h_comp := point.h
			if h_comp < height_line[x].h && h_comp < height_line[x+2].h && h_comp < heights[y][x+1].h && h_comp < heights[y+2][x+1].h {
				// Found a low point, now hunt the edges
				// Put this point in our new basin slice
				basins = append(basins, basin{basin_ix, 1})
				heights[y+1][x+1].basin_ix = basin_ix
				basin_ix++
			}
		}
	}

	// Found all low points, with a basin for each
	// Now exhaustively loop through the positions until all points with a height <9 have a basin
	var found_all_points = true

	for {
		found_all_points = true
		for y, line := range heights[1 : len(heights)-1] {
			for x, point := range line[1 : len(line)-1] {
				if point.h < 9 && point.basin_ix == -1 {
					// need a basin, haven't found one yet
					left := heights[y+1][x]
					right := heights[y+1][x+2]
					above := heights[y][x+1]
					below := heights[y+2][x+1]

					if same_basin(left, point) {
						heights[y+1][x+1].basin_ix = left.basin_ix
						basins[heights[y+1][x+1].basin_ix].size++
					} else if same_basin(right, point) {
						heights[y+1][x+1].basin_ix = right.basin_ix
						basins[heights[y+1][x+1].basin_ix].size++
					} else if same_basin(above, point) {
						heights[y+1][x+1].basin_ix = above.basin_ix
						basins[heights[y+1][x+1].basin_ix].size++
					} else if same_basin(below, point) {
						heights[y+1][x+1].basin_ix = below.basin_ix
						basins[heights[y+1][x+1].basin_ix].size++
					} else {
						found_all_points = false
					}
				}
			}
		}

		if found_all_points {
			break
		}
	}

	// now need to sort...
	var biggest_3 = [3]int{-1, -1, -1}

	for _, basin := range basins {
		// sort biggest_3
		if biggest_3[0] > biggest_3[1] {
			biggest_3 = [3]int{biggest_3[1], biggest_3[0], biggest_3[2]}
		}

		if biggest_3[1] > biggest_3[2] {
			biggest_3 = [3]int{biggest_3[0], biggest_3[2], biggest_3[1]}
		}

		if biggest_3[0] > biggest_3[1] {
			biggest_3 = [3]int{biggest_3[1], biggest_3[0], biggest_3[2]}
		}

		if basin.size > biggest_3[0] {
			biggest_3[0] = basin.size
		}
	}

	fmt.Println(biggest_3[0] * biggest_3[1] * biggest_3[2])

	return
}

func same_basin(compare_point point, this_point point) bool {
	return (compare_point.basin_ix != -1)
}

type point struct {
	x, y     int
	h        int
	basin_ix int
}

type basin struct {
	index int
	size  int
}
