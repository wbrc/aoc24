package main

import (
	"aoc/aoc"
	"io"
)

type dir func(x, y int) (int, int)

var dirs1 = []dir{
	func(x, y int) (int, int) { return x, y - 1 },
	func(x, y int) (int, int) { return x + 1, y - 1 },
	func(x, y int) (int, int) { return x + 1, y },
	func(x, y int) (int, int) { return x + 1, y + 1 },
	func(x, y int) (int, int) { return x, y + 1 },
	func(x, y int) (int, int) { return x - 1, y + 1 },
	func(x, y int) (int, int) { return x - 1, y },
	func(x, y int) (int, int) { return x - 1, y - 1 },
}

type grid [][]byte

func (g grid) at(x, y int) byte {
	if y < 0 || y >= len(g) {
		return 0
	}
	if x < 0 || x >= len(g[y]) {
		return 0
	}
	return g[y][x]
}

func (g grid) search(x, y int, search []byte, dirs []dir) int {
	num := 0
	for _, d := range dirs {
		xx, yy := x, y
		for i := 0; i < len(search); i++ {
			if g.at(xx, yy) != search[i] {
				break
			}
			xx, yy = d(xx, yy)
			if i == len(search)-1 {
				num++
			}
		}
	}
	return num
}

func One(in io.Reader) int {
	var g grid
	xes := []struct{ x, y int }{}
	for i, line := range aoc.Lines2(in) {
		l := []byte(line)
		for j, c := range l {
			if c == 'X' {
				xes = append(xes, struct{ x, y int }{j, i})
			}
		}
		g = append(g, l)
	}

	var num int
	for _, x := range xes {
		num += g.search(x.x, x.y, []byte("XMAS"), dirs1)
	}
	return num
}
