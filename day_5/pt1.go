// Run as `go run pt1.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type line struct {
	start      coordinate
	stop       coordinate
	horizontal bool
}

type coordinate struct {
	x, y int
}

type board struct {
	squares [][]int
}

func main() {
	line_instructions := os.Args[1:]
	regex := regexp.MustCompile("(?P<start_x>\\d+),(?P<start_y>\\d+) -> (?P<stop_x>\\d+),(?P<stop_y>\\d+)")
	groupNames := regex.SubexpNames()
	var lines = []line{}
	var max_x = 0
	var max_y = 0

	for _, instrct := range line_instructions {
		var start_x, start_y, stop_x, stop_y int
		match_sets := regex.FindAllStringSubmatch(instrct, -1)

		for _, match := range match_sets {
			for groupIdx, group := range match {
				int_val, _ := strconv.Atoi(group)
				name := groupNames[groupIdx]

				switch name {
				case "start_x":
					start_x = int_val
					if int_val > max_x {
						max_x = int_val
					}
				case "start_y":
					start_y = int_val
					if int_val > max_y {
						max_y = int_val
					}
				case "stop_x":
					stop_x = int_val
					if int_val > max_x {
						max_x = int_val
					}
				case "stop_y":
					stop_y = int_val
					if int_val > max_y {
						max_y = int_val
					}
				}
			}

			start := coordinate{start_x, start_y}
			stop := coordinate{stop_x, stop_y}

			if start_x == stop_x {
				lines = append(lines, line{start, stop, false})
			} else if start_y == stop_y {
				lines = append(lines, line{start, stop, true})
			}
		}
	}

	var my_board = board{make([][]int, max_x+1)}
	for x := range my_board.squares {
		my_board.squares[x] = make([]int, max_y+1)
		for y := range my_board.squares[x] {
			my_board.squares[x][y] = 0
		}
	}

	var multi_line_squares = 0

	for _, l := range lines {
		// min/maxing both works for now because only x ^ y will change
		xl, xh := min_max(l.start.x, l.stop.x)
		for x := xl; x <= xh; x++ {
			yl, yh := min_max(l.start.y, l.stop.y)
			for y := yl; y <= yh; y++ {
				my_board.squares[x][y]++
				if my_board.squares[x][y] == 2 {
					multi_line_squares++
				}
			}
		}
	}

	fmt.Println("Squares with multiple lines: ", multi_line_squares)

	// for _, r := range my_board.squares {
	// 	for _, c := range r {
	// 		if c > 9 {
	// 			fmt.Println("uhoh")
	// 			os.Exit(1)
	// 		}
	// 		if c == 0 {
	// 			fmt.Printf(".")
	// 		} else {
	// 			fmt.Printf("%d", c)
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return
}

func min_max(a, b int) (min, max int) {
	if a < b {
		return a, b
	}
	return b, a
}
