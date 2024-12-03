package day3

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

func TestSolve1(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	want := 161
	got := Solve1(input)
	assert.Equal(t, want, got)
}

func TestSolve2(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	want := 48
	got := Solve2(input)
	assert.Equal(t, want, got)
}
