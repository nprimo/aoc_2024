package day14

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`

func TestSolve1(t *testing.T) {
	want := 12
	got := Solve1(input)
	assert.Equal(t, want, got)
}

func TestMove(t *testing.T) {
	guard := Guard{
		pos: Pos{2, 4},
		vel: Pos{2, -3},
	}

	got := guard.Move(5, 11, 7).pos
	want := Pos{1, 3}
	assert.Equal(t, want, got)
}
