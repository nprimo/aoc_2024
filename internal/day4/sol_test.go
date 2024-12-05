package day4

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

func TestSolve1(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	got := Solve1(input)
	want := 18

	assert.Equal(t, want, got)
}
