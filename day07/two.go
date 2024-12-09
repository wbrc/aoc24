package main

import (
	"aoc/aoc"
	"io"
)

func powInt(base, exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}
	return res
}

func Two(in io.Reader) int {
	var sum int
	for line := range aoc.Lines(in) {
		expect := line.Split(":").IntAt(0)
		nums := line.Split(": ")[1:].Join("").Fields().Ints()
		for i := 0; i < powInt(3, len(nums)-1); i++ {
			ops := opsFromInt(i, len(nums)-1, 3)
			if calculate(ops, nums) == expect {
				sum += expect
				break
			}
		}
	}
	return sum
}
