package day12

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
`

func TestSolve1(t *testing.T) {
	got := Solve1(input)
	want := 1930
	assert.Equal(t, want, got)
}
