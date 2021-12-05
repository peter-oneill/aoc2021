package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type square struct {
	value int
	found bool
}

type line struct {
	values      [5]square
	found_count int
}

type board struct {
	rows          [5]line
	cols          [5]line
	remaining_sum int
	won           bool
}

func main() {
	number_order := strings.Split(os.Args[1], ",")
	fmt.Println(number_order)
	board_input := os.Args[2:]
	boards := []board{}
	var remaining_boards = 0

	for ii, arg := range board_input {
		val, _ := strconv.Atoi(arg)
		board_ix := ii / 25
		row_ix := (ii % 25) / 5
		col_ix := (ii % 25) % 5

		if len(boards) <= board_ix {
			var new_board = &board{[5]line{{[5]square{{0, false}}, 0}}, [5]line{{[5]square{{0, false}}, 0}}, 0, false}
			boards = append(boards, *new_board)
			remaining_boards++
		}

		this_board := &boards[board_ix]
		this_board.rows[row_ix].values[col_ix].value = val
		this_board.cols[col_ix].values[row_ix].value = val
		this_board.remaining_sum += val
	}

	var bingo = false

	for _, num := range number_order {
		num_val, _ := strconv.Atoi(num)

		for b_ix := range boards {
			b := &boards[b_ix]
			for r_ix := range b.rows {
				r := &b.rows[r_ix]
				for c_ix := range r.values {
					v := &r.values[c_ix]
					if v.value == num_val {
						v.found = true
						r.found_count++
						col := &b.cols[c_ix]
						col.values[r_ix].found = true
						col.found_count++
						b.remaining_sum -= num_val
						if !b.won && (r.found_count == 5 || col.found_count == 5) {
							fmt.Println("Board ", b_ix, " won!")
							b.won = true
							remaining_boards--
							if remaining_boards == 0 {
								fmt.Println("Found final board")
								bingo = true
								break
							}
						}
					}
				}

				if bingo {
					break
				}
			}

			if bingo {
				fmt.Println("Last winning board is ix ", b_ix)
				fmt.Println("Remaining sum: ", b.remaining_sum)
				fmt.Println("Final value found: ", num_val)
				fmt.Println("Score (remaining * last value): ", b.remaining_sum*num_val)
				break
			}
		}

		if bingo {
			break
		}
	}

	return
}
