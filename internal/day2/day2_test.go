package day2

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

func TestSolve1(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	want := 2
	got := Solve1(input)
	assert.Equal(t, want, got)
}

func TestSolve2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	want := 4
	got := Solve2(input)
	assert.Equal(t, want, got)
}
