package day7

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestSolve1(t *testing.T) {
	want := 3749
	got := Solve1(input)
	assert.Equal(t, want, got)
}

func TestSolve2(t *testing.T) {
	want := 11387
	got := Solve2(input)
	assert.Equal(t, want, got)
}

func TestPossibilities(t *testing.T) {
	nums := []int{6, 8, 6, 15}
	possibleResults3Ops([]int{nums[0]}, nums[1:])
}
