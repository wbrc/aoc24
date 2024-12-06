package main

import (
	"aoc/aoc"
	"io"
)

func hasCycle(obstacles aoc.Set[position], maxX, maxY int, pos position) bool {
	direction := position{0, -1}

	path := map[position]aoc.Set[position]{}

	for {
		if path[pos] == nil {
			path[pos] = aoc.NewSet[position]()
		}
		if d, ok := path[pos]; ok && d.Has(direction) {
			return true
		}
		path[pos].Set(direction)

		ahead := move(pos, direction)

		if ahead.x < 0 || ahead.x > maxX || ahead.y < 0 || ahead.y > maxY {
			break
		}
		if obstacles.Has(ahead) {
			direction = rotate90(direction)
			continue
		}

		pos = ahead
	}

	return false
}

func Two(in io.Reader) int {
	var maxX, maxY int
	var pos position
	obstacles := aoc.NewSet[position]()
	for y, line := range aoc.Lines2(in) {
		maxY = y
		for x, c := range line {
			maxX = x
			if c == '#' {
				obstacles.Set(position{x, y})
			} else if c == '^' {
				pos = position{x, y}
			}
		}
	}

	path := walk(obstacles, maxX, maxY, pos)
	path.Delete(pos)

	cycles := 0
	for p := range path {
		obstacles.Set(p)
		if hasCycle(obstacles, maxX, maxY, pos) {
			cycles++
		}
		obstacles.Delete(p)
	}

	return cycles
}
