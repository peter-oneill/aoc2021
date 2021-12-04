package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	binary_values := os.Args[1:]
	one_counts := []int{}

	// Run through the first input value as we don't yet know the length of the input vals
	for _, v := range binary_values[0] {
		var val = 0
		if v == '1' {
			val = 1
		}
		one_counts = append(one_counts, val)
	}

	// For every other input value, we can now just index directly
	for _, bin_number := range binary_values[1:] {
		for ii, digit := range bin_number {
			if digit == '1' {
				one_counts[ii]++
			}
		}
	}

	fmt.Println("one counts", one_counts)

	number_of_binary_values := len(binary_values)
	gamma_values := make([]string, number_of_binary_values)
	epsilon_values := make([]string, number_of_binary_values)

	// Find the most common value for each location
	for ii, count := range one_counts {
		if count > number_of_binary_values/2 {
			gamma_values[ii] = "1"
			epsilon_values[ii] = "0"
		} else {
			gamma_values[ii] = "0"
			epsilon_values[ii] = "1"
		}
	}

	gamma := binary_digit_string_arr_to_int(gamma_values)
	epsilon := binary_digit_string_arr_to_int(epsilon_values)

	fmt.Printf("gamm: %d, epsilon: %d\n", gamma, epsilon)
	fmt.Println("gamma * epsilon = ", gamma*epsilon)

	return
}

func binary_digit_string_arr_to_int(binary_array []string) (ret_val int64) {
	value_string := strings.Join(binary_array, "")
	ret_val, _ = strconv.ParseInt(value_string, 2, 0)
	return
}
