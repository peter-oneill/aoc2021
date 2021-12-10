// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(os.Args[1], ",")
	num_crabs := len(input)
	var positions = make([]int, num_crabs)
	var min, max int

	for ii, s := range input {
		v, _ := strconv.Atoi(s)
		positions[ii] = v

		if ii == 0 {
			min, max = v, v
		}

		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	fmt.Println("min, max: ", min, max)

	var best_cost int = -1

	var guess = (min + max) / 3
	var best_guess int = guess

	for {
		var current_cost int = 0
		for _, p := range positions {
			current_cost += fuel_cost(p, guess)
		}

		if best_cost == -1 {
			best_cost = current_cost
			best_guess = guess
		} else if current_cost < best_cost {
			if best_guess < guess-1 {
				min = best_guess
			} else if best_guess == guess-1 {
				min = guess
			} else if best_guess > guess+1 {
				max = best_guess
			} else if best_guess == guess+1 {
				max = guess
			}
			best_guess = guess
			best_cost = current_cost
		} else if current_cost > best_cost {
			// this cost is higher than our best
			if best_guess < guess-1 {
				max = guess
			} else if best_guess == guess-1 {
				max = best_guess
			} else if best_guess > guess+1 {
				min = guess
			} else if best_guess == guess+1 {
				min = best_guess
			}
		} else {
			// Got same cost - must have got to the same square.  Try one higher
			fmt.Println("At ", guess, " cost is ", current_cost, "min: ", min, "max: ", max)
			guess = guess + 1
			continue
		}

		fmt.Println("At ", guess, " cost is ", current_cost, "min: ", min, "max: ", max)
		if min == max {
			break
		}
		guess = (min + max) / 2

	}

	fmt.Println("Best cost: ", best_cost)

	return
}

func fuel_cost(position int, goal int) (cost int) {
	distance := int(math.Abs(float64(position - goal)))
	cost = (distance + 1) * distance / 2
	return
}
