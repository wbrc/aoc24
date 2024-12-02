package main

import (
	"aoc/aoc"
	"io"
)

func isSafe2(s []int) bool {
	ss := make([]int, len(s)-1)
	for i := range s {
		copy(ss, s[:i])
		copy(ss[i:], s[i+1:])
		if isSafe(ss) {
			return true
		}
	}
	return false
}
func Two(in io.Reader) int {
	res := 0
	for line := range aoc.Lines(in) {
		if isSafe2(line.Fields().Ints()) {
			res++
		}
	}
	return res
}
