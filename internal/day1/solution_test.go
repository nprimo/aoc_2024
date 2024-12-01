package day1

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

func TestSolve1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	want := 11
	got := Solve1(input)
	assert.Equal(t, want, got)
}

func TestSolve2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	want := 31
	got := Solve2(input)
	assert.Equal(t, want, got)
}
