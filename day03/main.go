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

var day = "day03"

var Example = []byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)

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
