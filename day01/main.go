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

var day = "day01"

var Example = []byte(`3   4
4   3
2   5
1   3
3   9
3   3
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
