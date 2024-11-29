package aoc

func Must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}
