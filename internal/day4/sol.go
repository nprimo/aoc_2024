package day4

import (
	"strings"
)

func Solve1(input string) int {
	// get a matrix

	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")

	matrix := make([][]string, len(rows))
	for i, row := range rows {
		matrix_r := strings.Split(row, "")
		matrix[i] = matrix_r
	}

	sum := 0
	for y, row := range matrix {
		for x := range row {
			for _, d := range dir {
				if isXMAS(Pos{x, y}, d, matrix) {
					sum++
				}
			}
		}
	}
	return sum
}

const xmas = "XMAS"

type Pos struct {
	X int
	Y int
}

func (p *Pos) Add(other Pos) {
	p.X += other.X
	p.Y += other.Y
}

var (
	up    = Pos{0, -1}
	down  = Pos{0, 1}
	right = Pos{1, 0}
	left  = Pos{-1, 0}
	rUp   = Pos{1, -1}
	rDown = Pos{1, 1}
	lUp   = Pos{-1, -1}
	lDown = Pos{-1, 1}
	dir   = []Pos{
		up,
		down,
		right,
		left,
		rUp,
		rDown,
		lUp,
		lDown,
	}
)

func isXMAS(curr Pos, dir Pos, matrix [][]string) bool {
	for i := range xmas {
		if curr.X < 0 || curr.X >= len(matrix[0]) {
			return false
		}
		if curr.Y < 0 || curr.Y >= len(matrix) {
			return false
		}
		if matrix[curr.Y][curr.X] != string(xmas[i]) {
			return false
		}
		curr.Add(dir)
	}
	return true
}
