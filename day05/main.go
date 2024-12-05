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

var day = "day05"

var Example = []byte(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
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
