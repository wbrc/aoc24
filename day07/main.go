package main

import (
	"aoc/aoc"
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"time"
)

var day = "day07"

var Example = []byte(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`)

func main() {
	part := flag.Int("p", 0, "part 1 or 2 (0 for both)")
	example := flag.Bool("x", false, "use example input")
	flag.Parse()

	data := Example
	if !*example {
		data = aoc.Must(os.ReadFile(path.Join(day, "input.txt")))
	}

	fmt.Printf("Puzzle %s\n", day)
	if *part == 0 || *part == 1 {
		start := time.Now()
		result := One(bytes.NewReader(data))
		fmt.Printf("Part 1 result: %d in %s\n", result, time.Since(start))
	}
	if *part == 0 || *part == 2 {
		start := time.Now()
		result := Two(bytes.NewReader(data))
		fmt.Printf("Part 2 result: %d in %s\n", result, time.Since(start))
	}
}
