package main

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####
`

func TestSolve1(t *testing.T) {
	want := 3
	got := Solve1(input)
	assert.Equal(t, want, got)
}
