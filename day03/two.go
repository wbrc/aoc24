package main

import (
	"aoc/aoc"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func Two(in io.Reader) int {
	raw := aoc.Must(io.ReadAll(in))
	input := strings.ReplaceAll(string(raw), "\n", "")
	muls := regexp.MustCompilePOSIX(`mul\([0-9]{1,3},[0-9]{1,3}\)|do(n't)?\(\)`).FindAllString(input, -1)
	res := 0
	enabled := true
	for _, mul := range muls {
		if strings.HasPrefix(mul, "do") {
			if mul == "do()" {
				enabled = true
			} else {
				enabled = false
			}
			continue
		}
		if !enabled {
			continue
		}
		var a, b int
		aoc.Must(fmt.Sscanf(mul, "mul(%d,%d)", &a, &b))
		res += a * b
	}
	return res
}
