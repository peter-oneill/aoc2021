package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	binary_values := os.Args[1:]

	oxygen_str := filter_binary_vals(binary_values, true, '1', 0)[0]
	co2_scrub_str := filter_binary_vals(binary_values, false, '0', 0)[0]

	oxygen, _ := strconv.ParseInt(oxygen_str, 2, 0)
	co2_scrub, _ := strconv.ParseInt(co2_scrub_str, 2, 0)

	fmt.Printf("oxygen: %d, co2_scrub: %d\n", oxygen, co2_scrub)
	fmt.Println("oxygen * co2_scrub = ", oxygen*co2_scrub)

	return
}

func filter_binary_vals(numbers []string, most_common bool, decider_digit rune, digit_ix int) (result []string) {
	var one_count = 0
	var zeroes = []string{}
	var ones = []string{}

	for _, number := range numbers {

		if number[digit_ix] == '1' {
			one_count++
			ones = append(ones, number)
		} else {
			zeroes = append(zeroes, number)
		}
	}

	var chosen = []string{}

	if one_count*2 == len(numbers) {
		if decider_digit == '1' {
			chosen = ones
		} else {
			chosen = zeroes
		}
	} else {
		ones_are_most_common := one_count*2 > len(numbers)

		if (most_common && ones_are_most_common) || (!most_common && !ones_are_most_common) {
			chosen = ones
		} else {
			chosen = zeroes
		}
	}

	if len(chosen) == 1 {
		return chosen
	}

	digit_ix++
	if digit_ix < len(chosen[0]) {
		chosen = filter_binary_vals(chosen, most_common, decider_digit, digit_ix)
	}

	return chosen
}
