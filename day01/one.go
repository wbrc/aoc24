package main

import (
	"aoc/aoc"
	"io"
	"sort"
)

func One(in io.Reader) int {
	var left, right []int
	for line := range aoc.Lines(in) {
		left = append(left, line.Fields().IntAt(0))
		right = append(right, line.Fields().IntAt(1))
	}
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		diff := right[i] - left[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum
}
