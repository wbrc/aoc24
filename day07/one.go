package main

import (
	"aoc/aoc"
	"io"
	"strconv"
)

type op byte

const (
	add op = iota
	mul
	concat
)

func (o op) do(a, b int) int {
	switch o {
	case add:
		return a + b
	case mul:
		return a * b
	case concat:
		var buf [32]byte
		str := strconv.AppendInt(buf[:0], int64(a), 10)
		str = strconv.AppendInt(str, int64(b), 10)
		return aoc.Must(strconv.Atoi(string(str)))
	}
	panic("invalid operation")
}

func opsFromInt(x int, length int, base int) []op {
	ops := []byte(strconv.FormatInt(int64(x), base))
	res := make([]op, length)
	start := length - len(ops)
	for i, o := range ops {
		res[start+i] = op(o - '0')
	}

	return res
}

func calculate(ops []op, nums []int) int {
	result := nums[0]
	for i, n := range nums[1:] {
		result = ops[i].do(result, n)
	}
	return result
}

func One(in io.Reader) int {
	var sum int
	for line := range aoc.Lines(in) {
		expect := line.Split(":").IntAt(0)
		nums := line.Split(": ")[1:].Join("").Fields().Ints()
		for i := 0; i < 1<<(len(nums)-1); i++ {
			ops := opsFromInt(i, len(nums)-1, 2)
			if calculate(ops, nums) == expect {
				sum += expect
				break
			}
		}
	}
	return sum
}
