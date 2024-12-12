package day11

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `125 17
`

func TestSolve1(t *testing.T) {
	want := 55312
	got := Solve1(input)
	assert.Equal(t, want, got)
}

// func TestSolve2(t *testing.T) {
// 	want := 55312
// 	got := Solve2(input)
// 	assert.Equal(t, want, got)
// }
