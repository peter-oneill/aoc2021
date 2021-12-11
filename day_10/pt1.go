// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input_lines := strings.Split(os.Args[1], "\n")
	var score = 0

	for _, line := range input_lines {
		var openers = []rune{'.'}
		var openers_ix = len(openers)
		var good_line = true

		for _, r := range []rune(line) {
			prev_ix := openers_ix - 1
			// fmt.Println("ix, previx, r, prevr", openers_ix, prev_ix, string(r), string(openers[prev_ix]))
			switch r {
			case '}':
				switch openers[prev_ix] {
				case '{':
					openers_ix--
				case '(', '<', '[':
					score += 1197
					good_line = false
					break
				}
			case ')':
				switch openers[prev_ix] {
				case '(':
					openers_ix--
				case '{', '<', '[':
					score += 3
					good_line = false
					break
				}
			case '>':
				switch openers[prev_ix] {
				case '<':
					openers_ix--
				case '(', '{', '[':
					score += 25137
					good_line = false
					break
				}
			case ']':
				switch openers[prev_ix] {
				case '[':
					openers_ix--
				case '(', '<', '{':
					score += 57
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
	}

	fmt.Println(score)

	return
}
