// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(os.Args[1], "\n")
	var sum int = 0

	for _, line := range lines {
		samples := strings.Split(strings.Split(line, " | ")[0], " ")
		var zero, one, two, three, four, five, six, seven, eight, nine []rune

		for _, sample := range samples {
			switch len(sample) {
			case 2:
				one = []rune(sample)
			case 3:
				seven = []rune(sample)
			case 4:
				four = []rune(sample)
			case 7:
				eight = []rune(sample)
			}
		}

		for _, sample := range samples {
			if len(sample) == 6 {
				rune_s := []rune(sample)
				if contains_segments(rune_s, four) {
					nine = rune_s
				} else if contains_segments(rune_s, one) {
					zero = rune_s
				} else {
					six = rune_s
				}
			}
		}

		for _, sample := range samples {
			if len(sample) == 5 {
				rune_s := []rune(sample)
				if contains_segments(rune_s, one) {
					three = rune_s
				} else if contains_segments(nine, rune_s) {
					five = rune_s
				} else {
					two = rune_s
				}
			}
		}

		one = sortRunes(one)
		two = sortRunes(two)
		three = sortRunes(three)
		four = sortRunes(four)
		five = sortRunes(five)
		six = sortRunes(six)
		seven = sortRunes(seven)
		eight = sortRunes(eight)
		nine = sortRunes(nine)
		zero = sortRunes(zero)

		var four_digit_num_slices = []rune{}

		for _, w := range strings.Split(strings.Split(line, " | ")[1], " ") {
			sorted_w := sortRunes([]rune(w))
			var digit rune

			if same_segments(sorted_w, zero) {
				digit = '0'
			} else if same_segments(sorted_w, one) {
				digit = '1'
			} else if same_segments(sorted_w, two) {
				digit = '2'
			} else if same_segments(sorted_w, three) {
				digit = '3'
			} else if same_segments(sorted_w, four) {
				digit = '4'
			} else if same_segments(sorted_w, five) {
				digit = '5'
			} else if same_segments(sorted_w, six) {
				digit = '6'
			} else if same_segments(sorted_w, seven) {
				digit = '7'
			} else if same_segments(sorted_w, eight) {
				digit = '8'
			} else if same_segments(sorted_w, nine) {
				digit = '9'
			}

			four_digit_num_slices = append(four_digit_num_slices, digit)
		}
		four_digit_num_string := string(four_digit_num_slices)
		four_digit_num, _ := strconv.Atoi(four_digit_num_string)
		sum += four_digit_num
	}

	fmt.Println("sum = ", sum)

	return
}

func contains_segments(holder []rune, segments []rune) bool {
	for _, test_seg := range segments {
		var found_loc = false
		for _, location := range holder {
			if test_seg == location {
				found_loc = true
				break
			}
		}
		if !found_loc {
			return false
		}
	}
	return true
}

func same_segments(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for ii := range a {
		if a[ii] != b[ii] {
			return false
		}
	}

	return true
}

func sortRunes(input []rune) []rune {
	sort.Sort(sortRuneSlice(input))
	return input
}

type sortRuneSlice []rune

func (s sortRuneSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneSlice) Len() int {
	return len(s)
}
