package main

import (
	"aoc/aoc"
	"io"
)

type position struct {
	x, y int
}

func move(p position, direction position) position {
	return position{p.x + direction.x, p.y + direction.y}
}

func rotate90(direction position) position {
	switch direction {
	case position{0, -1}:
		return position{1, 0}
	case position{1, 0}:
		return position{0, 1}
	case position{0, 1}:
		return position{-1, 0}
	case position{-1, 0}:
		return position{0, -1}
	}
	panic("invalid direction")
}

func walk(obstacles aoc.Set[position], maxX, maxY int, pos position) aoc.Set[position] {
	direction := position{0, -1}

	path := aoc.NewSet[position]()

	for {
		path.Set(pos)
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

	return path
}

func One(in io.Reader) int {
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

	return path.Len()
}
