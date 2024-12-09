package day9

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `2333133121414131402`

func TestSolve1(t *testing.T) {
	want := 1928
	got := Solve1(input)
	assert.Equal(t, want, got)

}
