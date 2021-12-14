// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"strings"
)

type instruction struct {
	in_pair   string
	out_pair1 string
	out_pair2 string
}

func build_initial_pairs(input string) map[string]int {
	var pairs = make(map[string]int)

	for ii := 0; ii < len(input)-1; ii++ {
		pair := input[ii : ii+2]
		count, exists := pairs[pair]
		if !exists {
			pairs[pair] = 1
		} else {
			pairs[pair] = count + 1
		}
	}
	return pairs
}

func build_instructions(input string) []instruction {
	var instructions []instruction
	for _, line := range strings.Split(input, "\n") {
		input_pair := line[0:2]
		insert_char := line[6:7]
		output_pair1 := input_pair[0:1] + insert_char
		output_pair2 := insert_char + input_pair[1:2]
		instructions = append(instructions, instruction{input_pair, output_pair1, output_pair2})
	}
	return instructions
}

func main() {
	input_sections := strings.Split(os.Args[1], "\n\n")

	var pairs = build_initial_pairs(input_sections[0])
	instructions := build_instructions(input_sections[1])

	for ii := 0; ii < 10; ii++ {
		var next_pairs = make(map[string]int)
		for _, instruct := range instructions {
			input_count, input_exists := pairs[instruct.in_pair]

			if input_exists {
				out1_count, out1_exists := next_pairs[instruct.out_pair1]
				if out1_exists {
					next_pairs[instruct.out_pair1] = out1_count + input_count
				} else {
					next_pairs[instruct.out_pair1] = input_count
				}

				out2_count, out2_exists := next_pairs[instruct.out_pair2]
				if out2_exists {
					next_pairs[instruct.out_pair2] = out2_count + input_count
				} else {
					next_pairs[instruct.out_pair2] = input_count
				}
			}
		}
		pairs = next_pairs
	}

	var min_count int = -1
	var max_count int = -1
	var char_count = make(map[string]int)

	for pair, pair_count := range pairs {
		count1, exists1 := char_count[pair[0:1]]
		if exists1 {
			char_count[pair[0:1]] = count1 + pair_count
		} else {
			char_count[pair[0:1]] = pair_count
		}

		count2, exists2 := char_count[pair[1:2]]
		if exists2 {
			char_count[pair[1:2]] = count2 + pair_count
		} else {
			char_count[pair[1:2]] = pair_count
		}
	}

	// This result includes dupes as every letter except first and last is counted twice (occurs at end of one pair, start of another)
	first_input_letter := input_sections[0][0:1]
	input_last_ix := len(input_sections[0]) - 1
	last_input_letter := input_sections[0][input_last_ix : input_last_ix+1]

	for letter, count := range char_count {
		if letter == first_input_letter {
			count -= 1
		}
		if letter == last_input_letter {
			count -= 1
		}
		count /= 2

		if letter == first_input_letter {
			count += 1
		}
		if letter == last_input_letter {
			count += 1
		}

		char_count[letter] = count
	}

	// find min, max
	for _, count := range char_count {
		if count > max_count {
			max_count = count
		} else if count < min_count || min_count == -1 {
			min_count = count
			fmt.Println()
		}
	}

	fmt.Println("max, min, diff: ", max_count, min_count, max_count-min_count)

	return
}
