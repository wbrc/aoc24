package main

import (
	"aoc/aoc"
	"io"
)

func (g grid) search2(x, y int) bool {
	c := 0

	if g.at(x-1, y-1) == 'M' && g.at(x+1, y+1) == 'S' {
		c++
	} else if g.at(x-1, y-1) == 'S' && g.at(x+1, y+1) == 'M' {
		c++
	}

	if g.at(x-1, y+1) == 'M' && g.at(x+1, y-1) == 'S' {
		c++
	} else if g.at(x-1, y+1) == 'S' && g.at(x+1, y-1) == 'M' {
		c++
	}

	if c == 2 {
		return true
	}
	return false
}

func Two(in io.Reader) int {
	var g grid
	xes := []struct{ x, y int }{}
	for i, line := range aoc.Lines2(in) {
		l := []byte(line)
		for j, c := range l {
			if c == 'A' {
				xes = append(xes, struct{ x, y int }{j, i})
			}
		}
		g = append(g, l)
	}

	var num int
	for _, x := range xes {
		if g.search2(x.x, x.y) {
			num++
		}
	}
	return num
}
