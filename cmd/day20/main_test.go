package main

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`

func TestSolve1(t *testing.T) {
	got := Solve1(input)
	want := 84
	assert.Equal(t, want, got)
}
