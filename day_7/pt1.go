// Run as `go run pt1.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(os.Args[1], ",")
	num_crabs := len(input)
	var positions = make([]int, num_crabs)

	for ii, s := range input {
		positions[ii], _ = strconv.Atoi(s)
	}

	sort.Ints(positions)
	median := positions[(num_crabs / 2)]
	var distance float64 = 0

	for _, p := range positions {
		distance += math.Abs(float64(p - median))
	}

	fmt.Println("Median: ", median, ", total distance: ", distance)

	return
}
