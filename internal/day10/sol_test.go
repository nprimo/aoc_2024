package day10

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestSolve1(t *testing.T) {
	got := Solve1(input)
	want := 36
	assert.Equal(t, want, got)
}
