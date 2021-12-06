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
	var today_fish []int
	var tomorrow_fish []int
	var num_fish int

	for _, f := range fish_strs {
		fish_num, _ := strconv.Atoi(f)
		today_fish = append(today_fish, fish_num)
		num_fish++
	}

	for day := 0; day < 80; day++ {
		for _, f := range today_fish {
			if f == 0 {
				tomorrow_fish = append(tomorrow_fish, 8, 6)
				num_fish += 2
			} else {
				tomorrow_fish = append(tomorrow_fish, f-1)
				num_fish++
			}
		}

		today_fish = tomorrow_fish
		tomorrow_fish = nil
		fmt.Println("Day ", day, " number of fish = ", num_fish, " len fish ", len(today_fish))
		num_fish = 0
	}

	return
}
