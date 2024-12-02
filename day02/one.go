package main

import (
	"aoc/aoc"
	"io"
)

func isSafe(s []int) bool {
	var inc, dec bool
	for i := 0; i < len(s)-1; i++ {
		diff := s[i] - s[i+1]
		if diff < 0 {
			inc = true
			diff = -diff
		} else if diff > 0 {
			dec = true
		}

		if inc && dec {
			return false
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func One(in io.Reader) int {
	res := 0
	for line := range aoc.Lines(in) {
		if isSafe(line.Fields().Ints()) {
			res++
		}
	}
	return res
}
