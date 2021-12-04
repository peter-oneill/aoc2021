package main

import (
    "os"
    "fmt"
    "strconv"
)

func main() {
    num_times_deeper := 0
    prev_depth, _ := strconv.Atoi(os.Args[1])

    for _, depth_str := range os.Args[2:] {
        depth, _ := strconv.Atoi(depth_str)
        if depth > prev_depth {
            num_times_deeper++
        }
        prev_depth = depth
    }

    fmt.Println("Number of times deeper:", num_times_deeper)
    return
}

