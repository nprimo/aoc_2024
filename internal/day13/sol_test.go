package day13

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
`

func TestSolve1(t *testing.T) {
	want := 480
	got := Solve1(input)
	assert.Equal(t, want, got)
}

func TestSolve2(t *testing.T) {
	Solve2(input)
}

func TestGetButton(t *testing.T) {
	testCases := []struct {
		instructions string
		want         [2]int
	}{

		{`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`,
			[2]int{80, 40},
		},
		{
			`Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`,
			[2]int{},
		},
	}

	for _, tc := range testCases {
		got := GetButtonComb(tc.instructions)
		assert.Equal(t, tc.want, got)

	}
}
