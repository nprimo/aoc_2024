package main

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
`

func TestSolve1(t *testing.T) {
	got := Solve1(input)
	want := 7036
	assert.Equal(t, want, got)
}
