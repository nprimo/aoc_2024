package day6

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

func TestSolve1(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`
	want := 41
	got := Solve1(input)
	assert.Equal(t, want, got)
}
