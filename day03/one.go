package main

import (
	"aoc/aoc"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func One(in io.Reader) int {
	raw := aoc.Must(io.ReadAll(in))
	input := strings.ReplaceAll(string(raw), "\n", "")
	muls := regexp.MustCompilePOSIX(`mul\([0-9]{1,3},[0-9]{1,3}\)`).FindAllString(input, -1)
	res := 0
	for _, mul := range muls {
		var a, b int
		aoc.Must(fmt.Sscanf(mul, "mul(%d,%d)", &a, &b))
		res += a * b
	}
	return res
}
