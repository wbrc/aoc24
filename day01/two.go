package main

import (
	"aoc/aoc"
	"io"
)

func Two(in io.Reader) int {
	var left []int
	right := map[int]int{}
	for line := range aoc.Lines(in) {
		left = append(left, line.Fields().IntAt(0))
		right[line.Fields().IntAt(1)]++
	}

	sum := 0
	for _, v := range left {
		sum += v * right[v]
	}

	return sum
}
