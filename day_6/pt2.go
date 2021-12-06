// Run as `go run pt1.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fish_strs := strings.Split(os.Args[1], ",")
	// var today_fish []int
	// var tomorrow_fish []int
	// var today_num_fish int
	// var tomorrow_num_fish int

	var pools = [9]int{0}

	for _, f := range fish_strs {
		fish_num, _ := strconv.Atoi(f)
		pools[fish_num]++
	}

	for day := 0; day < 256; day++ {
		old_pools := pools
		pools = [9]int{0}
		for p_ix, num_fish := range old_pools {
			if p_ix == 0 {
				pools[6] += num_fish
				pools[8] += num_fish
			} else {
				pools[p_ix-1] += num_fish
			}
		}

		total_fish := sum_pools(pools)
		fmt.Println("Day ", day, " number of fish = ", total_fish)
	}

	return
}

func sum_pools(pools [9]int) (sum int) {
	for _, val := range pools {
		sum += val
	}
	return
}
