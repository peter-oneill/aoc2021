// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	lines := strings.Split(os.Args[1], "\n")
	var digit_count = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

		for _, w := range strings.Split(strings.Split(line, " | ")[1], " ") {
			sorted_w := sortRunes([]rune(w))

			if same_segments(sorted_w, zero) {
				digit_count[0]++
			} else if same_segments(sorted_w, one) {
				digit_count[1]++
			} else if same_segments(sorted_w, two) {
				digit_count[2]++
			} else if same_segments(sorted_w, three) {
				digit_count[3]++
			} else if same_segments(sorted_w, four) {
				digit_count[4]++
			} else if same_segments(sorted_w, five) {
				digit_count[5]++
			} else if same_segments(sorted_w, six) {
				digit_count[6]++
			} else if same_segments(sorted_w, seven) {
				digit_count[7]++
			} else if same_segments(sorted_w, eight) {
				digit_count[8]++
			} else if same_segments(sorted_w, nine) {
				digit_count[9]++
			}
		}
	}

	fmt.Println(digit_count[0], digit_count[1], digit_count[2], digit_count[3], digit_count[4], digit_count[5], digit_count[6], digit_count[7], digit_count[8], digit_count[9])
	total_1478s := digit_count[1] + digit_count[4] + digit_count[7] + digit_count[8]
	fmt.Println("total 1s, 4s, 7s, 8s = ", total_1478s)

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
