package main

import (
    "os"
    "fmt"
    "strconv"
)

func main() {
    num_times_deeper := 0
    initial := 1
    offset := 3

    for ii, depth_str := range os.Args[initial + offset:] {
        prev_depth, _ := strconv.Atoi(os.Args[ii + initial])
        depth, _ := strconv.Atoi(depth_str)
        if depth > prev_depth {
            num_times_deeper++
        }
    }

    fmt.Println("Number of times deeper:", num_times_deeper)
    return
}

