package main

import (
	"aoc/aoc"
	"io"
	"sort"
)

func sortSeq(seq []int, after map[int]aoc.Set[int]) {
	sort.Slice(seq, func(i, j int) bool {
		if after, ok := after[seq[i]]; ok {
			if after.Has(seq[j]) {
				return true
			}
		}

		if after, ok := after[seq[j]]; ok {
			if after.Has(seq[i]) {
				return false
			}
		}

		panic("unreachable")
	})
}

func Two(in io.Reader) int {
	after := map[int]aoc.Set[int]{}
	parseOrder := true
	sum := 0

	for line := range aoc.Lines(in) {
		if line == "" {
			parseOrder = false
			continue
		}

		if parseOrder {
			var a, b int
			line.Scanf("%d|%d", &a, &b)
			if after[a] == nil {
				after[a] = aoc.NewSet[int]()
			}
			after[a].Set(b)
			continue
		}

		seq := line.Split(",").Ints()
		if isSorted(seq, after) {
			continue
		}

		sortSeq(seq, after)

		sum += seq[len(seq)/2]
	}

	return sum
}
