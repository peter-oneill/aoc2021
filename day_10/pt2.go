// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input_lines := strings.Split(os.Args[1], "\n")
	var scores []int

	for _, line := range input_lines {
		var openers = []rune{'.'}
		var openers_ix = len(openers)
		var good_line = true

		for _, r := range []rune(line) {
			prev_ix := openers_ix - 1
			switch r {
			case '}':
				switch openers[prev_ix] {
				case '{':
					openers_ix--
					openers = openers[:openers_ix]
				case '(', '<', '[':
					good_line = false
					break
				}
			case ')':
				switch openers[prev_ix] {
				case '(':
					openers_ix--
					openers = openers[:openers_ix]
				case '{', '<', '[':
					good_line = false
					break
				}
			case '>':
				switch openers[prev_ix] {
				case '<':
					openers_ix--
					openers = openers[:openers_ix]
				case '(', '{', '[':
					good_line = false
					break
				}
			case ']':
				switch openers[prev_ix] {
				case '[':
					openers_ix--
					openers = openers[:openers_ix]
				case '(', '<', '{':
					good_line = false
					break
				}
			default:
				if openers_ix >= len(openers) {
					openers = append(openers, r)
				} else {
					openers[openers_ix] = r
				}
				openers_ix++
			}

			if !good_line {
				break
			}
		}

		if good_line {
			var score = 0
			for ii := range openers {
				switch openers[len(openers)-ii-1] {
				case '(':
					score = 5*score + 1
				case '[':
					score = 5*score + 2
				case '{':
					score = 5*score + 3
				case '<':
					score = 5*score + 4
				}
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	middle := scores[len(scores)/2]

	fmt.Println(middle)

	return
}
