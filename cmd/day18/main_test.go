package main

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`

func TestSolve1(t *testing.T) {
	got := Solve1(input)
	want := 22
	assert.Equal(t, want, got)
}