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

func (l Line) Scanf(format string, a ...interface{}) {
	n := Must(fmt.Sscanf(string(l), format, a...))
	if n != len(a) {
		panic("invalid format")
	}
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

func Lines2(r io.Reader) iter.Seq2[int, Line] {
	s := bufio.NewScanner(r)
	return func(yield func(int, Line) bool) {
		i := 0
		for s.Scan() {
			if !yield(i, Line(s.Text())) {
				break
			}
			i++
		}
	}
}

type Set[A comparable] map[A]struct{}

func NewSet[A comparable]() Set[A] {
	return make(map[A]struct{})
}

func (s Set[A]) Set(a A) {
	s[a] = struct{}{}
}

func (s Set[A]) Has(a A) bool {
	_, ok := s[a]
	return ok
}

func (s Set[A]) Delete(a A) {
	delete(s, a)
}

func (s Set[A]) Len() int {
	return len(s)
}
