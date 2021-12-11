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
	var heights [][]int

	for y, input_line := range input_lines {
		if y == len(heights) {
			heights = append(heights, []int{})
		}

		heights[y] = append(heights[y], 10)

		for _, input_digit := range []rune(input_line) {
			height, _ := strconv.Atoi(string(input_digit))
			heights[y] = append(heights[y], height)
		}

		heights[y] = append(heights[y], 10)
	}

	var ten_line = make([]int, len(heights[0]))

	for x, _ := range ten_line {
		ten_line[x] = 10
	}

	heights = append(heights, ten_line)
	heights = append([][]int{ten_line}, heights...)

	var risk_sum int = 0

	for y, height_line := range heights[1 : len(heights)-1] {

		for x, height := range height_line[1 : len(height_line)-1] {
			if height < height_line[x] && height < height_line[x+2] && height < heights[y][x+1] && height < heights[y+2][x+1] {
				risk_sum += height + 1
			}
		}
	}

	fmt.Println(risk_sum)

	return
}
