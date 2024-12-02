package aoc

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"strconv"
	"strings"
)

func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

type Line string

func (l Line) Scanf(format string, a ...interface{}) (int, error) {
	return fmt.Sscanf(string(l), format, a...)
}

func (l Line) Fields() Fields {
	return strings.Fields(string(l))
}

func (l Line) Split(sep string) Fields {
	return strings.Split(string(l), sep)
}

type Fields []string

func (f Fields) StringAt(i int) string {
	return f[i]
}

func (f Fields) IntAt(i int) int {
	return Must(strconv.Atoi(f[i]))
}

func (f Fields) Ints() []int {
	res := make([]int, len(f))
	for i, s := range f {
		res[i] = Must(strconv.Atoi(s))
	}
	return res
}

func Lines(r io.Reader) iter.Seq[Line] {
	s := bufio.NewScanner(r)
	return func(yield func(Line) bool) {
		for s.Scan() {
			if !yield(Line(s.Text())) {
				break
			}
		}
	}
}
