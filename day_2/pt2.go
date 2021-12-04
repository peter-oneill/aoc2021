package main

import (
	"fmt"
	"os"
	"strconv"
)

type coordinate struct {
	x, y int
}

func main() {
	instructions := os.Args[1:]
	position := coordinate{0, 0}
	aim := 0

	for ii, instruction := range instructions {
		if ii%2 != 0 {
			continue
		}

		dir := instruction
		mag, _ := strconv.Atoi(instructions[ii+1])

		switch dir {
		case "forward":
			position.x += mag
			position.y += aim * mag
			break
		case "down":
			aim += mag
			break
		case "up":
			aim -= mag
			break
		}
	}

	fmt.Printf("x: %d, y: %d\n", position.x, position.y)
	fmt.Println("x * y = ", position.x*position.y)

	return
}
